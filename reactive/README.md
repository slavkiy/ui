# reactive

`reactive` - компактный Go-пакет для реактивного состояния, вычисляемых значений, событий, списков и batch-обновлений.

Все примеры ниже предполагают импорт:

```go
import (
    "fmt"
    "time"

    "github.com/slavkiy/ui/reactive"
)
```

## Что есть в пакете

- `Signal[T]` - текущее реактивное значение
- `Computed[T]` - вычисляемое значение из явных зависимостей
- `Event[T]` - transient-события
- `List[T]` - список с уведомлениями
- `Batch(fn)` - группировка обновлений
- `Memo[T]` - ленивый одноразовый вычисляемый кеш
- `Debounce`, `Throttle` - потоковые фильтры по времени
- `Storage`, `FileStorage`, `Persist` - сохранение состояния
- `Undoable[T]` - undo/redo вокруг сигнала
- `Animate` - плавная анимация `float64` сигнала
- `Subscription` - отменяемая подписка

## Быстрый пример

```go
count := reactive.NewSignal("count", 1)

count.Subscribe(func(value int) {
    fmt.Println("count:", value)
})

count.Set(2)
count.Update(func(v int) int { return v + 1 })
// count: 2
// count: 3
```

## Signal

`Signal[T]` хранит текущее значение и рассылает новые значения подписчикам.

### Создание

Используйте явный `scope` и имя сигнала. Это самый предсказуемый и безопасный вариант.

```go
name := reactive.NewSignalInScope[string]("app/internal", "name", "Slava")
fmt.Println(name.Get()) // Slava
```

`scope` — это просто строка пространства имён, например:

- `main`
- `app/internal`
- `app/admin`

Внутри глобального реестра это хранится как:

```go
app/internal::name
```

### Прикладное правило

Обычно пишут так:

```go
userName := reactive.NewSignalInScope[string]("app/internal", "userName", "Slava")
```

И получают так:

```go
loaded := reactive.GetSignalInScope[string]("app/internal", "userName")
fmt.Println(loaded.Get())
```

### Почему так лучше

Потому что это явно задаёт пространство имён и не зависит от структуры исходников, пакетов или файловой системы.

- одинаковый `name` в одном `scope` и том же `T` -> `panic`
- одинаковый `name` в одном `scope` и другом `T` -> допустимо
- одинаковый `name` в разных `scope` -> допустимо

Пример:

```go
titleText := reactive.NewSignalInScope[string]("app/internal", "title", "hello")
countText := reactive.NewSignalInScope[int]("app/internal", "title", 42)

fmt.Println(reactive.GetSignalInScope[string]("app/internal", "title").Get()) // hello
fmt.Println(reactive.GetSignalInScope[int]("app/internal", "title").Get())    // 42
```

Важная идея: путь указывает разработчик сам. Пакет не пытается угадывать его автоматически.

### Более короткий вариант

Если вы не хотите писать `scope` вручную, можно использовать локальный `main`-scope:

```go
count := reactive.NewSignalInScope[int]("main", "count", 1)
fmt.Println(reactive.GetSignalInScope[int]("main", "count").Get()) // 1
```

### `NewSignal` и `GetSignal`

Для простых случаев есть короткие варианты:

```go
name := reactive.NewSignal("name", "Slava")
fmt.Println(reactive.GetSignal[string]("name").Get())
```

Но в библиотеке/сервисах обычно предпочтительнее явно писать `scope`, чтобы не возникали коллизии имён между независимыми частями приложения.

### Пример конфликта

```go
// panic: duplicate signal registration for same scope/name/type
s1 := reactive.NewSignalInScope[string]("app/internal", "name", "Slava")
s2 := reactive.NewSignalInScope[string]("app/internal", "name", "Other")
```

### Тип подставляется через generics

Вы не пишете имя типа в строке, а просто указываете `T` в общем вызове:

```go
v := reactive.GetSignalInScope[string]("app/internal", "title")
```

Это и есть смысл подстановки типа: `GetSignalInScope[T](scope, name)`.

Для обычного кода обычно лучше передавать `*Signal[T]` явно, а `GetSignal` использовать только для глобального registry и именованных state-объектов.

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

- `Set(value)` - заменить значение
- `Update(fn)` - получить текущее значение и заменить результатом функции
- `Update(nil)` - ничего не делает

### Подписки

У `Signal` есть несколько видов подписок.

#### `Subscribe(func(T))`

```go
status := reactive.NewSignal("status", "idle")
sub := status.Subscribe(func(value string) {
    fmt.Println("status:", value)
})
defer sub.Unsubscribe()

status.Set("loading")
```

#### `Effect(func(T))`

```go
logSub := status.Effect(func(value string) {
    fmt.Println("log:", value)
})
defer logSub.Unsubscribe()
```

`Effect` подходит для побочных действий: логирование, отправка метрик, синхронизация с внешним сервисом.

#### `SubscribeEffect(func())`

```go
refreshSub := status.SubscribeEffect(func() {
    fmt.Println("refresh")
})
defer refreshSub.Unsubscribe()
```

Используйте, когда важно только событие изменения, а значение можно не читать.

#### `SubscribeChan(bufferSize)`

```go
updates := reactive.NewSignal("updates", 0)
ch, sub := updates.SubscribeChan(4)
defer sub.Unsubscribe()

go func() {
    for v := range ch {
        fmt.Println("from channel:", v)
    }
}()

updates.Set(1)
updates.Set(2)
```

Отправка в channel не блокирует поток навсегда; при заполнении буфера новые значения просто пропускаются. `Unsubscribe` закрывает канал.

## Computed

`Computed[T]` вычисляет значение из явных зависимостей.

```go
first := reactive.NewSignal("first", 2)
second := reactive.NewSignal("second", 3)

sum := reactive.NewComputed(func() int {
    return first.Get() + second.Get()
}, first, second)

fmt.Println(sum.Get()) // 5

first.Set(10)
fmt.Println(sum.Get()) // 13
```

Ключевая идея: зависимости указываются явно. Если вы забыли передать `second`, пересчёт не произойдёт при изменении `second`.

```go
func NewComputed[T any](compute func() T, dependencies ...Dependency) *Computed[T]
func (c *Computed[T]) Get() T
func (c *Computed[T]) Subscribe(fn func(T)) *Subscription
func (c *Computed[T]) SubscribeEffect(fn func()) *Subscription
```

Используйте `Computed`, когда значение явно является производным от других реактивных данных.

## Event

`Event[T]` не хранит текущее состояние. Он просто эмитит значения в подписчиков.

```go
clicked := reactive.NewEvent[string]()
sub := clicked.Subscribe(func(value string) {
    fmt.Println("clicked:", value)
})
defer sub.Unsubscribe()

clicked.Emit("save")
clicked.Emit("cancel")
```

Используйте `Event` для кликов, команд, сообщений, одноразовых уведомлений.

```go
func NewEvent[T any]() *Event[T]
func (e *Event[T]) Emit(value T)
func (e *Event[T]) Subscribe(fn func(T)) *Subscription
```

## List

`List[T]` представляет динамический набор значений и уведомляет подписчиков после изменений.

```go
items := reactive.NewList[int]()
items.Append(1, 2, 3)

items.Subscribe(func(v []int) {
    fmt.Println("items:", v)
})

items.Set([]int{10, 20})
items.Clear()
```

```go
func NewList[T any]() *List[T]
func (l *List[T]) Get() []T
func (l *List[T]) Set(items []T)
func (l *List[T]) Append(items ...T)
func (l *List[T]) Clear()
func (l *List[T]) Subscribe(fn func([]T)) *Subscription
```

`Get()` возвращает копию среза, поэтому внешний код не может изменить внутреннее состояние напрямую.

## Batch

`Batch` откладывает уведомления до конца внешнего блока.

```go
first := reactive.NewSignal("first", 1)
second := reactive.NewSignal("second", 2)

reactive.Batch(func() {
    first.Set(10)
    reactive.Batch(func() {
        second.Set(20)
    })
})
```

Используйте `Batch`, когда нужно изменить несколько связанных сигналов как одно логическое обновление.

## Memo

`Memo[T]` вычисляет значение лениво и кеширует результат после первого вызова.

```go
memo := reactive.NewMemo(func() int {
    fmt.Println("compute")
    return 2 * 21
})

fmt.Println(memo.Get()) // compute -> 42
fmt.Println(memo.Get()) // 42
```

Хорошо подходит для дорогих вычислений, которые не должны повторяться.

## Debounce и Throttle

### Debounce

`Debounce` ждёт паузу после серии обновлений и отправляет последнее значение.

```go
query := reactive.NewSignal("query", "")
settled := reactive.Debounce(query, 300*time.Millisecond)

settled.Subscribe(func(value string) {
    fmt.Println("search:", value)
})

query.Set("g")
query.Set("go")
query.Set("golang")
```

Используйте, когда много быстрых изменений должны приводить к одному финальному действию.

### Throttle

`Throttle` ограничивает частоту событий.

```go
scroll := reactive.NewSignal("scroll", 0)
limited := reactive.Throttle(scroll, 100*time.Millisecond)

limited.Subscribe(func(value int) {
    fmt.Println("scroll:", value)
})

scroll.Set(10)
scroll.Set(20)
```

Подходит для scroll/mouse-move/частых UI-событий, где важна частота, а не каждое промежуточное значение.

## Storage и Persist

### Storage

`Storage` - контракт для хранения пар ключ-значение.

```go
type Storage interface {
    Get(key string, value any) error
    Set(key string, value any) error
    Delete(key string) error
}
```

### FileStorage

```go
storage := reactive.NewFileStorage("settings.json")
var theme string
if err := storage.Get("theme", &theme); err == nil {
    fmt.Println("theme:", theme)
}

_ = storage.Set("theme", "dark")
```

`FileStorage` хранит JSON-данные в файле.

### Persist

```go
storage := reactive.NewFileStorage("settings.json")
theme := reactive.Persist(storage, "theme", "light")

fmt.Println(theme.Get()) // light

theme.Set("dark")
```

`Persist` создаёт `Signal`, заполняет её из storage и сохраняет последующие значения автоматически.

## Undoable

`Undoable[T]` оборачивает `Signal[T]` и хранит историю изменений.

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
```

```go
func NewUndoable[T any](signal *Signal[T]) *Undoable[T]
func (u *Undoable[T]) Get() T
func (u *Undoable[T]) Set(value T)
func (u *Undoable[T]) Undo() bool
func (u *Undoable[T]) Redo() bool
func (u *Undoable[T]) Signal() *Signal[T]
```

## Animation

`Animate` двигает `Signal[float64]` от текущего значения к target в течение заданного времени.

```go
progress := reactive.NewSignal("progress", 0.0)
sub := reactive.Animate(progress, 100, time.Second, reactive.EaseOut)
defer sub.Unsubscribe()
```

Доступные easing-функции:

- `Linear`
- `EaseIn`
- `EaseOut`
- `EaseInOut`

Если `easing == nil`, используется `Linear`.

## Subscription

`Subscription` - безопасная отменяемая подписка.

```go
sub := signal.Subscribe(func(v int) {
    fmt.Println(v)
})

sub.Unsubscribe()
sub.Unsubscribe() // безопасно, повторный вызов игнорируется
```

```go
type Subscription struct {
    // internal
}

func (s *Subscription) Unsubscribe()
```

## Когда что использовать

- `Signal` - текущее состояние, которое нужно читать и обновлять
- `Computed` - производное значение из явных зависимостей
- `Event` - одноразовые события или команды
- `List` - динамическая коллекция с уведомлениями
- `Batch` - несколько обновлений как одна логическая операция
- `Debounce` - ждать, пока пользователь закончит ввод
- `Throttle` - ограничить частоту событий
- `Persist` - синхронизация состояния с хранилищем
- `Undoable` - отмена/повтор действий
- `Animate` - плавное изменение значения по времени

## Пара замечаний по модели

- подписчики вызываются вне mutex
- внутри `Subscribe`/`Effect` лучше избегать тяжёлых блокировок и долгих операций
- для многопоточного использования use-case: одни goroutines пишут в signal, другие читают через `.Get()` или подписки
- `Computed` не пытается автоматически отслеживать все `Get()` вызовы; зависимости должны быть переданы явно

## Итог

`reactive` - это маленькая, предсказуемая реактивная система под Go:

- простой `Signal`
- явные зависимости для `Computed`
- много полезных преложений вокруг `Signal`
- безопасный базовый контракт для UI-состояния и событий
