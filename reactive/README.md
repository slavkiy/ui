# reactive

`reactive` - небольшой generic-пакет для хранения состояния, подписок, производных значений, событий и типовых потоков обновлений.

Все примеры предполагают такой импорт:

```go
import (
    "fmt"
    "time"

    "github.com/slavkiy/ui/reactive"
)
```

## Быстрый пример

```go
count := reactive.NewSignal("count", 1)

stop := count.Subscribe(func(value int) {
    fmt.Println("count:", value)
})
defer stop.Unsubscribe()

count.Set(2)          // count: 2
count.Update(func(v int) int { return v + 1 }) // count: 3
```

`Signal[T]` хранит текущее значение типа `T`. `Set` и `Update` уведомляют callbacks после изменения. Подписки нужно отменять через `Unsubscribe`, когда владелец состояния больше не используется.

## Signal

### Создание и поиск

```go
name := reactive.NewSignal("name", "Slava")
fmt.Println(name.Get()) // Slava

sameName := reactive.GetSignal[string]("name")
fmt.Println(sameName.Get()) // Slava
```

`NewSignal(name, value)` создаёт сигнал и регистрирует его по имени. `GetSignal[T](name)` ищет зарегистрированный сигнал и возвращает `nil`, если имя отсутствует или тип не совпадает. Пустое имя допустимо, но такие сигналы не следует искать через registry.

Глобальный registry нужен, когда сигнал создаётся в одном месте, а получается по имени в другом. Для обычного кода лучше передавать `*Signal[T]` явно.

### Чтение, запись и обновление

```go
score := reactive.NewSignal("score", 10)

current := score.Get()
score.Set(20)
score.Update(func(value int) int {
    return value + 5
})

fmt.Println(current, score.Get()) // 10 25
```

Используйте `Set`, когда новое значение уже рассчитано. Используйте `Update`, когда новое значение зависит от текущего. `Update(nil)` ничего не делает.

## Подписки

У `Signal` есть три разных способа получать изменения.

### `Subscribe(func(T))`

Используйте, когда callback нужен для каждого обновления и вам нужно новое значение.

```go
signal := reactive.NewSignal("status", "idle")

subscription := signal.Subscribe(func(value string) {
    fmt.Println("render status:", value)
})
defer subscription.Unsubscribe()

signal.Set("loading")
```

### `Effect(func(T))`

`Effect` работает почти так же, как `Subscribe`, но явно обозначает побочное действие: запись в лог, отправку метрики, запрос или синхронизацию с внешней системой.

```go
stopLogging := signal.Effect(func(value string) {
    fmt.Println("log status:", value)
})
defer stopLogging.Unsubscribe()
```

Выбирайте `Effect`, когда callback не обновляет представление, а выполняет действие вне reactive-системы.

### `SubscribeEffect(func())`

Используйте, когда факт изменения важен, но новое значение не нужно.

```go
refresh := signal.SubscribeEffect(func() {
    fmt.Println("refresh requested")
})
defer refresh.Unsubscribe()
```

Для `Signal` `SubscribeEffect` является удобной обёрткой над `Effect`, а для `Computed` вызывается после пересчёта.

### `SubscribeChan(uint)`

Используйте канал, если изменения удобнее обрабатывать в goroutine или через `select`.

```go
updates := reactive.NewSignal("updates", 0)
values, subscription := updates.SubscribeChan(2)
defer subscription.Unsubscribe()

go func() {
    for value := range values {
        fmt.Println("received:", value)
    }
}()

updates.Set(1)
updates.Set(2)
```

Отправка в канал неблокирующая. Если буфер заполнен, новое обновление пропускается. `Unsubscribe` удаляет подписку и закрывает канал.

Все callbacks вызываются вне mutex. Не держите внешние mutex вокруг операций, если callback может вызвать другой reactive-метод.

## Computed

`Computed[T]` хранит значение, которое вычисляется из других reactive-объектов. Зависимости передаются явно в `NewComputed`.

```go
first := reactive.NewSignal("first", 2)
second := reactive.NewSignal("second", 3)

sum := reactive.NewComputed(
    func() int {
        return first.Get() + second.Get()
    },
    first, second,
)

fmt.Println(sum.Get()) // 5
sum.Subscribe(func(value int) {
    fmt.Println("sum:", value)
})

first.Set(10) // sum: 13
```

Пакет не отслеживает автоматически вызовы `Get()` внутри функции. Если забыть передать `second`, `sum` не пересчитается после изменения `second`. Если зависимости не переданы вообще, вычисление выполнится только при создании.

`Computed` нужен для производных данных. Для побочного действия используйте `Subscribe` или `SubscribeEffect`:

```go
stop := sum.SubscribeEffect(func() {
    fmt.Println("sum was recomputed")
})
defer stop.Unsubscribe()
```

## Batch

`Batch` откладывает callbacks до завершения внешнего batch. Вложенные batch поддерживаются.

```go
first := reactive.NewSignal("first", 1)
second := reactive.NewSignal("second", 2)

first.SubscribeEffect(func() { fmt.Println("first changed") })
second.SubscribeEffect(func() { fmt.Println("second changed") })

reactive.Batch(func() {
    first.Set(10)
    reactive.Batch(func() {
        second.Set(20)
    })
    fmt.Println("inside batch")
})
// Сначала печатается "inside batch", затем callbacks обновлений.
```

Используйте `Batch`, когда несколько связанных значений меняются одной операцией и промежуточные состояния не должны обрабатываться UI или бизнес-логикой.

## Event

`Event[T]` не хранит текущее значение. Он передаёт только значения, отправленные через `Emit`.

```go
clicked := reactive.NewEvent[string]()

stop := clicked.Subscribe(func(button string) {
    fmt.Println("clicked:", button)
})
defer stop.Unsubscribe()

clicked.Emit("save")
clicked.Emit("cancel")
```

Используйте `Event` для кликов, команд, уведомлений и других transient-событий. Используйте `Signal`, если новое состояние должно быть доступно через `Get`.

## Memo

`Memo[T]` вычисляет значение лениво и только один раз.

```go
memo := reactive.NewMemo(func() int {
    fmt.Println("compute")
    return 2 * 21
})

fmt.Println(memo.Get()) // печатает compute, затем 42
fmt.Println(memo.Get()) // печатает только 42
```

Используйте `Memo`, когда вычисление дорогое, а его результат неизменен в течение жизни memo. `Memo` не зависит от `Signal` и не пересчитывается после изменений.

## Debounce и Throttle

### Debounce

`Debounce` ждёт период без новых обновлений и передаёт последнее значение.

```go
query := reactive.NewSignal("query", "")
settled := reactive.Debounce(query, 300*time.Millisecond)

settled.Subscribe(func(value string) {
    fmt.Println("run search for:", value)
})

query.Set("g")
query.Set("go")
query.Set("golang")
// При достаточно быстрых Set поиск запускается только для "golang".
```

Используйте для поиска, автосохранения и resize. Это уменьшает количество реакций во время серии быстрых изменений.

### Throttle

`Throttle` пропускает не более одного значения за указанный период.

```go
scroll := reactive.NewSignal("scroll", 0)
limited := reactive.Throttle(scroll, 100*time.Millisecond)

limited.Subscribe(func(value int) {
    fmt.Println("update position:", value)
})

scroll.Set(10)
scroll.Set(20) // может быть пропущено, если прошло меньше 100 ms
```

Используйте для scroll, mouse move и других частых потоков, когда регулярное ограничение частоты важнее последнего значения.

## Storage и Persist

### Storage

`Storage` - интерфейс для чтения, записи и удаления значений по ключу. Можно использовать собственную реализацию, например поверх базы данных или memory map.

```go
type MemoryStorage struct {
    values map[string]any
}

// Реализация Storage должна определить Get, Set и Delete.
```

Методы интерфейса:

- `Get(key, &value)` загружает значение в переданный указатель.
- `Set(key, value)` сохраняет значение.
- `Delete(key)` удаляет значение.

### FileStorage

`NewFileStorage(path)` создаёт JSON-хранилище и загружает существующий файл.

```go
storage := reactive.NewFileStorage("settings.json")

var theme string
if err := storage.Get("theme", &theme); err == nil {
    fmt.Println("theme:", theme)
}

_ = storage.Set("theme", "dark")
_ = storage.Delete("old-theme")
```

`Get` возвращает `os.ErrNotExist`, если ключ отсутствует. Ошибки `Set` и `Delete` нужно обрабатывать в приложении; в примерах выше они сокращены до `_`.

### Persist

`Persist` создаёт `Signal`, загружая начальное значение из `Storage`, и сохраняет каждое последующее изменение.

```go
storage := reactive.NewFileStorage("settings.json")
theme := reactive.Persist(storage, "theme", "light")

fmt.Println(theme.Get()) // сохранённая тема или light
theme.Set("dark")        // новое значение будет сохранено
```

Используйте `Persist` для пользовательских настроек и небольшого состояния между запусками. Для `Persist` с nil storage создаётся обычный signal без сохранения.

## Undoable

`Undoable[T]` оборачивает `Signal[T]` и хранит историю значений.

```go
text := reactive.NewSignal("text", "one")
editor := reactive.NewUndoable(text)

editor.Set("two")
editor.Set("three")

fmt.Println(editor.Get()) // three
editor.Undo()
fmt.Println(editor.Get()) // two
editor.Redo()
fmt.Println(editor.Get()) // three

signal := editor.Signal()
signal.Subscribe(func(value string) {
    fmt.Println("editor changed:", value)
})
```

- `Get` возвращает текущее значение wrapped signal.
- `Set` меняет значение и добавляет его в историю.
- `Undo` возвращает `false`, если отменять нечего.
- `Redo` возвращает `false`, если повторять нечего.
- `Signal` возвращает исходный signal для обычных подписок.

После `Undo` новый `Set` начинает новую ветку истории и очищает возможность `Redo`.

## Animation

`Animate` постепенно меняет `Signal[float64]` и возвращает `Subscription`, которой можно остановить анимацию.

```go
progress := reactive.NewSignal("progress", 0.0)
progress.Subscribe(func(value float64) {
    fmt.Printf("progress: %.2f\n", value)
})

animation := reactive.Animate(
    progress,
    100,
    time.Second,
    reactive.EaseOut,
)

time.Sleep(200 * time.Millisecond)
animation.Unsubscribe() // отменить анимацию досрочно
```

Доступные easing-функции:

- `Linear` - равномерное движение.
- `EaseIn` - медленный старт и ускорение.
- `EaseOut` - быстрый старт и замедление.
- `EaseInOut` - ускорение в начале и замедление в конце.

Если easing равен nil, используется `Linear`. Если duration меньше или равен нулю, target устанавливается сразу.

## List

`List[T]` хранит срез и отправляет подписчикам копию элементов после `Set`, `Append` или `Clear`.

```go
items := reactive.NewList[string]()

stop := items.Subscribe(func(value []string) {
    fmt.Println("items:", value)
})
defer stop.Unsubscribe()

items.Append("one", "two")
items.Set([]string{"first", "second"})
fmt.Println(items.Get())
items.Clear()
```

`Get` возвращает копию среза, поэтому изменение полученного результата не меняет внутренний список. Используйте `List`, когда нужна коллекция с уведомлениями; для одного значения используйте `Signal[[]T]`.

## Выбор инструмента

| Задача | Используйте |
| --- | --- |
| Хранить текущее состояние | `Signal[T]` |
| Получить сигнал по имени | `GetSignal[T]` |
| Выполнить производное вычисление | `Computed[T]` |
| Выполнить побочное действие | `Effect` |
| Реагировать без чтения нового значения | `SubscribeEffect` |
| Получать обновления через goroutine | `SubscribeChan` |
| Передать transient-событие | `Event[T]` |
| Объединить несколько изменений | `Batch` |
| Вычислить один раз лениво | `Memo[T]` |
| Подождать паузу после серии изменений | `Debounce` |
| Ограничить частоту изменений | `Throttle` |
| Сохранить состояние между запусками | `Persist` и `Storage` |
| Добавить Undo/Redo | `Undoable[T]` |
| Плавно изменить число | `Animate` |
| Хранить список с подписками | `List[T]` |

## Жизненный цикл подписки

Каждый метод подписки возвращает `*Subscription`:

```go
subscription := signal.Subscribe(func(value int) {
    fmt.Println(value)
})

subscription.Unsubscribe()
subscription.Unsubscribe() // безопасно: callback удаляется только один раз
```

Nil callback не регистрируется. Отменяйте подписки при остановке компонента, закрытии страницы или завершении фоновой задачи, чтобы не удерживать ненужные объекты.
