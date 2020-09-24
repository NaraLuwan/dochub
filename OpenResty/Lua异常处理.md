## pcall & xpcall

lua有两个函数可用于捕获异常： pcall 和 xpcall，这两个函数很类似，都会在保护模式下执行函数，效果类似try-catch，可捕获并处理异常。

两个函数的原型如下：
```lua
pcall (func [, arg1, ···])
xpcall (func, errfunc [, arg1, ···])
```

对比两个函数，xpcall多了一个异常处理函数参数 errfunc。对于pcall，异常处理完时只简单记录错误信息，然后释放调用栈空间，而对于xpcall，这个参数可用于在调用栈释放前跟踪到这些数据。效果如下：

```text
> f=function(...) error(...) end
> pcall(f, 123)
false   stdin:1: 123
> xpcall(f, function(e) print(debug.traceback()) return e end, 123)
stack traceback:
        stdin:1: in function <stdin:1>
        [C]: in function 'error'
        stdin:1: in function 'f'
        [C]: in function 'xpcall'
        stdin:1: in main chunk
        [C]: in ?
false   stdin:1: 123
```

值得注意的是， errfunc的传入参数是异常数据，函数结束时必须将这个数据返回，才能实现和 pcall 一样的返回值。

## 原理

pcall以一种"保护模式"来调用第一个参数，因此pcall可以捕获函数执行中的任何错误。

通常在错误发生时，希望落得更多的调试信息，而不只是发生错误的位置。但pcall返回时，它已经销毁了调用桟的部分内容。

Lua提供了xpcall函数，xpcall接收第二个参数——一个错误处理函数，当错误发生时，Lua会在调用桟展开（unwind）前调用错误处理函数，于是就可以在这个函数中使用debug库来获取关于错误的额外信息了。

debug库提供了两个通用的错误处理函数:

- debug.debug：提供一个Lua提示符，让用户来检查错误的原因
- debug.traceback：根据调用桟来构建一个扩展的错误消息

## Demo
```lua
function traceback( msg )  
    print("----------------------------------------")  
    print("LUA ERROR: " .. tostring(msg) .. "\n")  
    -- print detail err msg
    print(debug.traceback())  
    print("----------------------------------------")  
end  
  
local function main(arg1, arg2)  
    -- ...
    print(arg1)  
end  
  
xpcall(main, traceback) 
```
