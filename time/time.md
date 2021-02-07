# time

## Time

表示时间的结构体，由时间和地区组成

```
type Time struct {
    wall uint64
    ext  int64
    loc *Location
}
```

#### (t Time) Format(layout string) string

包内提供了预定义的layout格式  
手动传入格式也必须是2006年1月2日 下午3点4分5秒，因为内部实现是通过这个时间点的数字来判断格式的

```
const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)
```

和java的时间格式定义对照

|java|go|说明|
|---|---|---|
|yyyy|2006|年|
|yy|06|年|
|MM|01|月|
|dd|02|日|
|HH|15|小时|
|mm|04|分钟|
|ss|05|秒|
|SSS|.000|毫秒|

### Now() Time

返回当前的本地时间

#### (t Time) Location() *Location

Location返回t的地点和时区信息。

#### (t Time) Zone() (name string, offset int)

Zone计算t所在的时区，返回该时区的规范名（如"CET"）和该时区相对于UTC的时间偏移量（单位秒）。

#### (t Time) IsZero() bool

IsZero报告t是否代表Time零值的时间点，January 1, year 1, 00:00:00 UTC。

#### (t Time) Local() Time

Local返回采用本地和本地时区，但指向同一时间点的Time。

#### (t Time) UTC() Time

UTC返回采用UTC和零时区，但指向同一时间点的Time。

#### (t Time) In(loc *Location) Time

In返回采用loc指定的地点和时区，但指向同一时间点的Time。如果loc为nil会panic。

#### (t Time) Unix() int64

Unix将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位秒）。

#### (t Time) UnixNano() int64

UnixNano将t表示为Unix时间，即从时间点January 1, 1970
UTC到时间点t所经过的时间（单位纳秒）。如果纳秒为单位的unix时间超出了int64能表示的范围，结果是未定义的。注意这就意味着Time零值调用UnixNano方法的话，结果是未定义的。

#### (t Time) Equal(u Time) bool

判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息。

#### (t Time) Before(u Time) bool

如果t代表的时间点在u之前，返回真；否则返回假。

#### (t Time) After(u Time) bool

如果t代表的时间点在u之后，返回真；否则返回假。

#### (t Time) Date() (year int, month Month, day int)

返回时间点t对应的年、月、日。

#### (t Time) Clock() (hour, min, sec int)

返回t对应的那一天的时、分、秒。

#### (t Time) Year() int

返回时间点t对应的年份。

#### (t Time) Month() Month

返回时间点t对应那一年的第几月。

#### (t Time) ISOWeek() (year, week int)

返回时间点t对应的ISO 9601标准下的年份和星期编号。星期编号范围[1,53]，1月1号到1月3号可能属于上一年的最后一周，12月29号到12月31号可能属于下一年的第一周。

#### (t Time) YearDay() int

返回时间点t对应的那一年的第几天，平年的返回值范围[1,365]，闰年[1,366]。

#### (t Time) Day() int

返回时间点t对应那一月的第几日。

#### (t Time) Weekday() Weekday

返回时间点t对应的那一周的周几。

#### (t Time) Hour() int

返回t对应的那一天的第几小时，范围[0, 23]。

#### (t Time) Minute() int

返回t对应的那一小时的第几分种，范围[0, 59]。

#### (t Time) Second() int

返回t对应的那一分钟的第几秒，范围[0, 59]。

#### (t Time) Nanosecond() int

返回t对应的那一秒内的纳秒偏移量，范围[0, 999999999]。

#### (t Time) Add(d Duration) Time

Add返回时间点t+d。

#### (t Time) AddDate(years int, months int, days int) Time

AddDate返回增加了给出的年份、月份和天数的时间点Time。例如，时间点January 1, 2011调用AddDate(-1, 2, 3)会返回March 4, 2010。

AddDate会将结果规范化，类似Date函数的做法。因此，举个例子，给时间点October 31添加一个月，会生成时间点December 1。（从时间点November 31规范化而来）

#### (t Time) Sub(u Time) Duration

返回一个时间段t-u。如果结果超出了Duration可以表示的最大值/最小值，将返回最大值/最小值。要获取时间点t-d（d为Duration），可以使用t.Add(-d)。

#### (t Time) Round(d Duration) Time

返回距离t最近的时间点，该时间点应该满足从Time零值到该时间点的时间段能整除d；如果有两个满足要求的时间点，距离t相同，会向上舍入；如果d <= 0，会返回t的拷贝。

#### (t Time) Truncate(d Duration) Time

类似Round，但是返回的是最接近但早于t的时间点；如果d <= 0，会返回t的拷贝。

### Parse(layout, value string) (Time, error)

根据给定的格式解析字符串返回时间

parse一样使用规定的格式进行时间的格式化

```
parse, err := time.Parse("1/2/2006 15:04:05.000000000", "12/8/2015 1:00:00.120000123")
if err != nil {
    log.Fatal(err)
}
fmt.Println(parse) 
// 输出: 2015-12-08 01:00:00.120000123 +0000 UTC
format := parse.Format("01 02 2006 Jan Mon 15 04 05 .000 ")
// 输出: 12 08 2015 Dec Tue 01 00 00 .120 
fmt.Println(format)
```

### Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time

根据参数返回对应时区的时间

### Unix(sec int64, nsec int64) Time

根据秒数，纳秒数返回Unix时间

## Duration

表示时间段的结构体

```
type Duration int64
```

最小单位是nanosecond纳秒,能表达的最长时间是290年

#### Constants 常量

```
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)
```

整数乘以Duration后的类型还是Duration

```
//Duration
second := 10 * time.Second 
```

整数变量乘以Duration编译报错，原因是不同类型无法运算

```
value :=10
//invalid operation: value * time.Second (mismatched types int and time.Duration)
second := value * time.Second
```

必须将变量转换为Duration才能通过编译

```
value := 10
second := time.Duration(value) * time.Second
```

可以直接用10*Duration是因为golang自动帮我们转换了类型

#### (d Duration) Hours() float64

将时间段表示为float64类型的小时数。

#### (d Duration) Minutes() float64

将时间段表示为float64类型的分钟数。

#### (d Duration) Seconds() float64

将时间段表示为float64类型的秒数。

#### (d Duration) Nanoseconds() int64

将时间段表示为float64类型的纳秒数。

#### (d Duration) String() string

返回时间段采用"72h3m0.5s"格式的字符串表示。最前面可以有符号，数字+单位为一个单元，开始部分的0值单元会被省略；如果时间段<1s，会使用"ms"、"us"、"ns"来保证第一个单元的数字不是0；如果时间段为0，会返回"0"。

### ParseDuration(s string) (Duration, error)

ParseDuration解析一个时间段字符串。一个时间段字符串是一个序列，每个片段包含可选的正负号、十进制数、可选的小数部分和单位后缀，如"300ms"、"-1.5h"、"2h45m"。合法的单位有"ns"、"us" /"µs"、"ms"
、"s"、"m"、"h"。

### Since(t Time) Duration

返回从t到现在经过的时间，等价于time.Now().Sub(t)。

### Sleep(d Duration)

使当前goroutine睡眠至少(d Duration)的时间  
传入一个负数或0会立即返回。

### After(d Duration) <-chan Time

返回一个只接受Time的channel

After会在另一线程经过时间段d后向返回值发送当时的时间。等价于NewTimer(d).C。

### Tick(d Duration) <-chan Time

Tick是NewTicker的封装，只提供对Ticker的通道的访问。如果不需要关闭Ticker，本函数就很方便。

## Weekday

表示星期的结构体

```
type Weekday int
```

```
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
```

#### (d Weekday) String() string

返回Weekday对应的英文名（"Sunday"、"Monday"，……）

## Month

表示月份的结构体

```
type Month int
```

```
const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
```

#### (m Month) String() string

返回Month对应的英文名（"January"，"February"，……）

## Location

表示地区的结构体

**UTC 零时区**

**Local 系统本地**

### LoadLocation(name string) (*Location, error)

根据给定的名字返回对应时区的Location

如果name是""或"UTC"，返回UTC；如果name是"Local"，返回Local；否则name应该是IANA时区数据库里有记录的地点名（该数据库记录了地点和对应的时区），如"America/New_York"。

## Timer

C类型为只接受Time类型的channel

```
type Timer struct {
    C <-chan Time
    // ...
}
```

Timer类型代表单次时间事件。当Timer到期时，当时的时间会被发送给C，除非Timer是被AfterFunc函数创建的。

#### NewTimer(d Duration) *Timer

NewTimer创建一个Timer，它会在最少过去时间段d后到期，向其自身的C字段发送当时的时间。

#### AfterFunc(d Duration, f func()) *Timer

AfterFunc另起一个go程等待时间段d过去，然后调用f。它返回一个Timer，可以通过调用其Stop方法来取消等待和对f的调用。

#### (t *Timer) Reset(d Duration) bool

Reset使t重新开始计时，（本方法返回后再）等待时间段d过去后到期。如果调用时t还在等待中会返回真；如果t已经到期或者被停止了会返回假。

#### (t *Timer) Stop() bool

Stop停止Timer的执行。如果停止了t会返回真；如果t已经被停止或者过期了会返回假。Stop不会关闭通道t.C，以避免从该通道的读取不正确的成功。

## Ticker

Ticker保管一个通道，并每隔一段时间(Duration)向其传递"tick"。

```
type Ticker struct {
    C <-chan Time // 周期性传递时间信息的通道
    // ...
}
```

### NewTicker(d Duration) *Ticker

NewTicker返回一个新的Ticker，该Ticker包含一个通道字段，并会每隔时间段d就向该通道发送当时的时间。它会调整时间间隔或者丢弃tick信息以适应反应慢的接收者。如果d<=0会panic。关闭该Ticker可以释放相关资源。

### (t *Ticker) Stop()

Stop关闭一个Ticker。在关闭后，将不会发送更多的tick信息。Stop不会关闭通道t.C，以避免从该通道的读取不正确的成功。


