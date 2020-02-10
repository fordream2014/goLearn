
* 变量
    + 全局变量 首字母大写为外部全局变量，首字母小写为内部全局变量
    + 常量 关键字const，全局常量大小写规则同变量
* 分支与循环
    + switch 支持变量值匹配，表达式匹配
        ```
        // 值匹配
        func prize1(score int) string {
            switch score / 10 {
            case 0, 1, 2, 3, 4, 5:
                return "差"
            case 6, 7:
                return "及格"
            case 8:
                return "良"
            default:
                return "优"
            }
        }
        
        // 表达式匹配
        func prize2(score int) string {
            // 注意 switch 后面什么也没有
            switch {
                case score < 60:
                    return "差"
                case score < 80:
                    return "及格"
                case score < 90:
                    return "良"
                default:
                    return "优"
            }
        }
        ```
    + 不支持三元操作符
* 数组
    + Go语言会对数组下标越界做编译器检查，但是如果下标用变量表示，Go会在代码中插入下标越界检查的逻辑。
    + 相同类型，相同长度的数组之间可以相互赋值
    + 数组的赋值，浅拷贝。赋值的两个数组变量的值不会共享
* 切片
    + 
    