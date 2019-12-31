
## Context使用规则

    a）勿将Context作为struct的字段使用，而是对每个使用其的函数分别作参数使用，其需定义为函数或方法的第一个参数，一般叫作ctx；
    b）勿对Context参数传nil，未想好的使用那个Context，请传context.TODO；
    c）使用context传值仅可用作请求域的数据，其它类型数据请不要滥用；
    d）同一个Context可以传给使用其的多个goroutine，且Context可被多个goroutine同时安全访问

## 最佳实践:

    context.Background 只应用在最高等级，作为所有派生 context 的根。
    context.TODO 应用在不确定要使用什么的地方，或者当前函数以后会更新以便使用 context。
    context 取消是建议性的，这些函数可能需要一些时间来清理和退出。
    context.Value 应该很少使用，它不应该被用来传递可选参数。这使得 API 隐式的并且可以引起错误。取而代之的是，这些值应该作为参数传递。
    不要将 context 存储在结构中，在函数中显式传递它们，最好是作为第一个参数。
    永远不要传递不存在的 context 。相反，如果您不确定使用什么，使用一个 ToDo context。
    Context 结构没有取消方法，因为只有派生 context 的函数才应该取消 context。

## context 包允许以下方式创建和获得 context

    1. context.Background() Context
    这个函数返回一个空 context。这只能用于高等级（在 main 或顶级请求处理中）
    
    2. context.TODO() Context
    这个函数也是创建一个空 context。也只能用于高等级或当您不确定使用什么 context，或函数以后会更新以便接收一个 context 。
    这意味您（或维护者）计划将来要添加 context 到函数。
    
    3. context.WithValue(parent Context, key, val interface{}) (ctx Context, cancel CancelFunc)
    此函数接收 context 并返回派生 context，其中值 val 与 key 关联，并通过 context 树与 context 一起传递。
    这意味着一旦获得带有值的 context，从中派生的任何 context 都会获得此值。
    不建议使用 context 值传递关键参数，而是函数应接收签名中的那些值，使其显式化。
    
    4. context.WithCancel(parent Context) (ctx Context, cancel CancelFunc)
    返回派生 context 和取消函数。只有创建它的函数才能调用取消函数来取消此 context。
    如果您愿意，可以传递取消函数，但是，强烈建议不要这样做。
    这可能导致取消函数的调用者没有意识到取消 context 的下游影响。可能存在源自此的其他 context，这可能导致程序以意外的方式运行。
    简而言之，永远不要传递取消函数。
    
    5. context.WithDeadline(parent Context, d time.Time) (ctx Context, cancel CancelFunc)
    此函数返回其父项的派生 context，当截止日期超过或取消函数被调用时，该 context 将被取消。
    例如，
        您可以创建一个将在以后的某个时间自动取消的 context，并在子函数中传递它。
        当因为截止日期耗尽而取消该 context 时，获此 context 的所有函数都会收到通知去停止运行并返回。
    
    6. context.WithTimeout(parent Context, timeout time.Duration) (ctx Context, cancel CancelFunc)
    此函数类似于 context.WithDeadline。不同之处在于它将持续时间作为参数输入而不是时间对象。
    此函数返回派生 context，如果调用取消函数或超出超时持续时间，则会取消该派生 context。


