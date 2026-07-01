# trit — довідник

Повний довідник пакета `trit`: ментальна модель, стани й перетворення, логічні
оператори, помічники дефолтного стану, серіалізація, зрізові агрегати, таблиці
істинності й практичні рецепти.

Англійська версія: **[DOC.md](DOC.md)**.

## Зміст

- [Ментальна модель](#ментальна-модель)
- [Стани й обмеження Logicable](#стани-й-обмеження-logicable)
- [Огляд і встановлення Trit](#огляд-і-встановлення-trit)
- [Помічники дефолтного стану](#помічники-дефолтного-стану)
- [Перетворення](#перетворення)
- [Логічні оператори](#логічні-оператори)
- [Порівняння й впорядкування](#порівняння-й-впорядкування)
- [Парсинг](#парсинг)
- [Серіалізація](#серіалізація)
- [Зрізові агрегати](#зрізові-агрегати)
- [Помилки](#помилки)
- [Таблиці істинності](#таблиці-істинності)
- [Рецепти й поради](#рецепти-й-поради)

## Ментальна модель

`trit` реалізує трійкову (тернарну) логіку. `Trit` — це `int8` із трьома
значущими станами:

| Стан | Значення |
|------|----------|
| `False`   | будь-яке від'ємне число (канонічно `-1`) |
| `Unknown` | `0` |
| `True`    | будь-яке додатне число (канонічно `1`) |

Ключова властивість: **нульове значення — це `Unknown`**, тож неініціалізований
`Trit` уже осмислений. Саме це робить `trit` корисним усюди, де важить стан
«можливо» чи «не задано» — злиття конфігів, часткові оновлення, nullable-колонки
БД, логічні схеми — де звичайний `bool` не відрізнить «false» від «не задано».

Оскільки будь-яке від'ємне/додатне ціле рахується як False/True, нормалізуйте до
канонічних `-1/0/1` через `Norm`/`Val`, коли потрібне точне значення.

```go
import "github.com/goloop/trit/v2"
```

## Стани й обмеження Logicable

Багато функцій узагальнені над `Logicable` — набором типів, що можуть заступити
тернарне значення: `bool`, цілі види й сам `Trit`. Це дозволяє передавати `bool`,
`int` чи `Trit` взаємозамінно:

```go
trit.Define(true)  // True
trit.Define(0)     // Unknown
trit.Define(-5)    // False
```

`Define(v)` конвертує будь-який `Logicable` у відповідний `Trit`.

## Огляд і встановлення Trit

```go
func (t Trit) IsTrue() bool
func (t Trit) IsFalse() bool
func (t Trit) IsUnknown() bool
func (t *Trit) Set(v int) Trit
func (t Trit) Val() Trit    // нормалізована копія (-1/0/1)
func (t *Trit) Norm() Trit  // нормалізувати на місці
func (t Trit) Int() int
func (t Trit) String() string
```

Є також узагальнені пакетні предикати, що приймають будь-який `Logicable`:

```go
func IsTrue[T Logicable](t T) bool
func IsFalse[T Logicable](t T) bool
func IsUnknown[T Logicable](t T) bool
```

```go
t := trit.True
t.IsTrue()   // true
t.Int()      // 1
t.String()   // "True"
```

## Помічники дефолтного стану

Вони розв'язують стан `Unknown`; це й є причина братися за `trit`:

```go
func (t *Trit) Default(trit Trit) Trit
func (t *Trit) TrueIfUnknown() Trit
func (t *Trit) FalseIfUnknown() Trit
func (t *Trit) Clean() Trit

func Default[T Logicable](t *Trit, v T) Trit
```

`Default` виставляє значення лише якщо воно наразі `Unknown`, інакше лишає
недоторканим — ідеально для «застосувати дефолт, не перезаписуючи явний вибір».
`TrueIfUnknown`/`FalseIfUnknown` згортають `Unknown` до певного стану; `Clean`
скидає до `Unknown`.

```go
t := trit.Unknown
t.Default(trit.True) // тепер True

trit.Default(&cfg.Enabled, true) // виставити дефолт, лише якщо не задано
```

## Перетворення

```go
func (t Trit) CanBeBool() bool
func (t Trit) ToBool() (bool, error)
```

`ToBool` повертає булеве значення або [`ErrUnknownValue`](#помилки) для
`Unknown` — тож перетворення нерозв'язаного стану на `bool` є явною, перевірною
операцією, а не тихим здогадом. `CanBeBool` повідомляє, чи вдалося б
перетворення.

```go
b, err := trit.True.ToBool() // true, nil
_, err = trit.Unknown.ToBool() // errors.Is(err, trit.ErrUnknownValue)
```

## Логічні оператори

Кожен оператор — це метод, що бере `Trit` і повертає `Trit`, за тернарними
[таблицями істинності](#таблиці-істинності):

| Метод | Операція |
|-------|----------|
| `Not()`         | логічне НЕ |
| `And`, `Or`, `Xor` | І / АБО / виключне АБО |
| `Nand`, `Nor`, `Nxor` | заперечені І / АБО / XOR |
| `Imp`, `Nimp`   | імплікація (Лукасевич) та її заперечення |
| `Eq`, `Neq`     | еквівалентність (тоді й лише тоді) та її заперечення |
| `Min`, `Max`    | мінімум / максимум (False < Unknown < True) |
| `Ma`, `La`, `Ia` | абсорбція Modus-Ponens / Law-of / Implication |

```go
t1, t2 := trit.True, trit.False
t1.And(t2) // False
t1.Or(t2)  // True
t1.Xor(t2) // True
t1.Not()   // False
```

## Порівняння й впорядкування

```go
func (t Trit) Compare(o Trit) int
```

`Compare` впорядковує `False < Unknown < True` і повертає `-1`, `0` чи `1` — той
самий контракт, що й `cmp.Compare`, тож він прямо під'єднується до
`slices.SortFunc`.

## Парсинг

```go
func ParseTrit(s string) (Trit, error)
```

Парсить текстове представлення в `Trit`, повертаючи [`ErrInvalidTrit`](#помилки)
для нерозпізнаного рядка. Приймає поширені написання трьох станів (напр. `"true"`,
`"false"`, `"unknown"` та їхні короткі/числові форми).

## Серіалізація

`Trit` реалізує стандартні інтерфейси серіалізації, відображаючи `Unknown` на
«null» кожного формату:

| Інтерфейс | Unknown стає |
|-----------|--------------|
| `json.Marshaler`/`Unmarshaler` (`MarshalJSON`/`UnmarshalJSON`) | JSON `null` |
| `encoding.TextMarshaler`/`TextUnmarshaler` (`MarshalText`/`UnmarshalText`) | текстова форма |
| `database/sql` (`Value`/`Scan`) | SQL `NULL` |

```go
data, _ := json.Marshal(trit.Unknown) // null
var t trit.Trit
_ = json.Unmarshal([]byte("true"), &t) // True
```

Це робить `Trit` природним для nullable-булевої колонки: `NULL` ⇄ `Unknown`, а
два певні стани повертаються туди-назад.

## Зрізові агрегати

Зводять багато `Logicable`-значень (чи `iter.Seq`) до одного `Trit`:

```go
func All[T Logicable](t ...T) Trit        func AllSeq[T Logicable](seq iter.Seq[T]) Trit
func Any[T Logicable](t ...T) Trit        func AnySeq[T Logicable](seq iter.Seq[T]) Trit
func None[T Logicable](t ...T) Trit       func NoneSeq[T Logicable](seq iter.Seq[T]) Trit
func Known[T Logicable](ts ...T) Trit     func KnownSeq[T Logicable](seq iter.Seq[T]) Trit
func Consensus[T Logicable](trits ...T) Trit
func Majority[T Logicable](trits ...T) Trit
```

- `All` / `Any` / `None` — тернарні квантори.
- `Known` — чи кожне значення певне (не `Unknown`).
- `Consensus` — спільне значення, коли всі згодні, інакше `Unknown`.
- `Majority` — значення, яке має більш ніж половина.

```go
trit.All(true, true, trit.Unknown) // Unknown
trit.Any(false, trit.True)         // True
trit.Majority(true, true, false)   // True
```

## Помилки

```go
var ErrInvalidTrit  = errors.New("invalid trit value")
var ErrUnknownValue = errors.New("cannot convert Unknown to bool")
```

`ParseTrit` повертає `ErrInvalidTrit`; `ToBool` повертає `ErrUnknownValue` для
`Unknown`. Звіряйте будь-яку через `errors.Is`.

## Таблиці істинності

```
 Truth Tables of Three-valued logic
 (T=True, N=Unknown, F=False)

  A  | NA      A  | MA      A  | LA      A  | IA
 ----+----    ----+----    ----+----    ----+----
  F  |  T      F  |  F      F  |  F      F  |  F
  U  |  U      U  |  T      U  |  F      U  |  T
  T  |  F      T  |  T      T  |  T      T  |  F

  A | B | AND       A | B |  OR       A | B | XOR
 ---+---+------    ---+---+------    ---+---+------
  F | F |  F        F | F |  F        F | F |  F
  F | U |  F        F | U |  U        F | U |  U
  F | T |  F        F | T |  T        F | T |  T
  U | F |  F        U | F |  U        U | F |  U
  U | U |  U        U | U |  U        U | U |  U
  U | T |  U        U | T |  T        U | T |  U
  T | F |  F        T | F |  T        T | F |  T
  T | U |  U        T | U |  T        T | U |  U
  T | T |  T        T | T |  T        T | T |  F

  A | B | NAND      A | B | NOR       A | B | NXOR
 ---+---+------    ---+---+------    ---+---+------
  F | F |  T        F | F |  T        F | F |  T
  F | U |  T        F | U |  U        F | U |  U
  F | T |  T        F | T |  F        F | T |  F
  U | F |  T        U | F |  U        U | F |  U
  U | U |  U        U | U |  U        U | U |  U
  U | T |  U        U | T |  F        U | T |  U
  T | F |  T        T | F |  F        T | F |  F
  T | U |  U        T | U |  F        T | U |  U
  T | T |  F        T | T |  F        T | T |  T

  A | B | IMP       A | B |  EQ       A | B | MIN
 ---+---+------    ---+---+------    ---+---+------
  F | F |  T        F | F |  T        F | F |  F
  F | U |  T        F | U |  U        F | U |  F
  F | T |  T        F | T |  F        F | T |  F
  U | F |  U        U | F |  U        U | F |  F
  U | U |  T        U | U |  U        U | U |  U
  U | T |  T        U | T |  U        U | T |  U
  T | F |  F        T | F |  F        T | F |  F
  T | U |  U        T | U |  U        T | U |  U
  T | T |  T        T | T |  T        T | T |  T

  A | B | NIMP      A | B | NEQ       A | B | MAX
 ---+---+------    ---+---+------    ---+---+------
  F | F |  F        F | F |  F        F | F |  F
  F | U |  F        F | U |  U        F | U |  U
  F | T |  F        F | T |  T        F | T |  T
  U | F |  U        U | F |  U        U | F |  U
  U | U |  F        U | U |  U        U | U |  U
  U | T |  F        U | T |  U        U | T |  T
  T | F |  T        T | F |  T        T | F |  T
  T | U |  U        T | U |  U        T | U |  T
  T | T |  F        T | T |  F        T | T |  T
```

## Рецепти й поради

**Часткові оновлення конфігу.** Дайте опційним прапорам тип `trit.Trit`. Незадане
поле лишається `Unknown`, тож можна застосувати `trit.Default(&field, fallback)`,
не затираючи явний `False` — проблему, яку звичайний `bool` не розв'язує.

**Nullable-булеві.** Зберігайте `Trit` у nullable SQL-колонці: `NULL` ⇄ `Unknown`
через `Value`/`Scan`, а JSON `null` ⇄ `Unknown` через JSON-методи.

**Агрегуйте голоси.** Використовуйте `Consensus`, коли кожен вхід має погодитись,
`Majority` для простого голосування й `Known`, щоб перевірити, чи всі входи
вирішені, перед дією.

**Нормалізуйте перед порівнянням сирих значень.** Будь-яке від'ємне/додатне ціле
є False/True; викличте `Norm`/`Val`, щоб отримати канонічні `-1/0/1`, коли важить
точний `int8`.

**Конвертуйте свідомо.** Використовуйте `ToBool` (і перевіряйте помилку), а не
припускайте, що `Unknown` — це false; явна `ErrUnknownValue` тримає неоднозначність
видимою.
