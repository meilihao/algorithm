// https://github.com/QMHTMY/RustBook/blob/main/code/chapter03/infix_to_postfix.rs
// 前缀表达式要求所有运算符在处理的两个操作数之前，后缀表达式则要求其操作符在相应的操作数之后. 它们均不包含调整优先级的括号.
// 只有中缀才需要括号，前缀和后缀表达式的操作顺序完全由操作符的顺序决定, 所以不用括号
// 计算前缀表达式: 需要两个栈, 一个存操作符, 一个存操作数, 直接按照从左到右的顺序将其分别入栈即可. 计算时先将操作符号出栈，然后将两个操作数出栈，此时用操作符计算这两个操作数，结果再入栈. 接着再重复这个计算步骤，直到操作符号栈空，此时弹出操作数栈顶数据，这个值就是整个表达式的计算结果
// 计算后缀表达式: 仅用一个栈

/*
中缀表达式转前后缀表达式的方法:
1. 采用完全括号表达式

	得到完全括号表达式本身就很困难，而且移动字符再删除字符涉及修改字符串，所以这种方法还不够实用(或通用)
1. 其他
	
	步骤:
	1. 创建一个名为 op_stack 的空栈以保存运算符, 给输出创建一个空列表 postfix
	2. 通过使用字符串方法拆分将输入的中缀字符串转换为标记列表 src_str
	3. 从左到右扫描标记列表
	
		如果标记是操作数，将其附加到输出列表的末尾
		如果标记是左括号，将其压到 op_stack 上
		如果标记是右括号，则弹出 op_stack，直到删除到相应左括号，将此过程中弹出的运算符加入 postfix
		如果标记是运算符 + - * /，则压入 op_stack, 但要先弹出 op_stack 中更高或相等优先级的运算符到 postfix
	4. 当输入处理完后，检查 op_stack, 仍在栈上的运算符都可弹出到 postfix
*/
use std::collections::HashMap;

#[derive(Debug)]
struct Stack<T> {
    top: usize,
    data: Vec<T>,
}

impl<T> Stack<T> {
    fn new() -> Self {
        Stack {
            top: 0,
            data: Vec::new()
        }
    }

    fn push(&mut self, val: T) {
        self.data.push(val);
        self.top += 1;
    }

    fn pop(&mut self) -> Option<T> {
        if self.top == 0 { return None; }
        self.top -= 1;
        self.data.pop()
    }

    fn peek(&self) -> Option<&T> {
        if self.top == 0 { return None; }
        self.data.get(self.top - 1)
    }

    fn is_empty(&self) -> bool {
        0 == self.top
    }
}

// 同时检测多种括号
fn par_match(open: char, close: char) -> bool {
    let opens = "([{";
    let closers = ")]}";
    opens.find(open) == closers.find(close)
}

// 检测括号是否匹配
fn par_checker3(infix: &str) -> bool {
    let mut char_list = Vec::new();
    for c in infix.chars() {
        char_list.push(c);
    }

    let mut index = 0;
    let mut balance = true;
    let mut stack = Stack::new();
    while index < char_list.len() && balance {
        let c = char_list[index];
        if '(' == c || '[' == c || '{' == c {
            stack.push(c);
        }

        if ')' == c || ']' == c || '}' == c {
            if stack.is_empty() {
                balance = false;
            } else {
                let top = stack.pop().unwrap();
                if !par_match(top, c) {
                    balance = false;
                }
            }
        }

        index += 1;
    }

    balance && stack.is_empty()
}

fn infix_to_postfix(infix: &str) -> Option<String> {
    // 括号匹配检验
    if !par_checker3(infix) { return None; }

    // 设置各个符号的优先级
    let mut prec = HashMap::new();
    prec.insert("(", 1); prec.insert(")", 1); // 左括号赋予最低的优先级，这样，与其进行比较的任何运算符都具有更高的优先级
    prec.insert("+", 2); prec.insert("-", 2);
    prec.insert("*", 3); prec.insert("/", 3);

    // ops 保存操作符号、postfix 保存后缀表达式
    let mut op_stack = Stack::new();
    let mut postfix = Vec::new();
    for token in infix.split_whitespace() {
        if ("A" <= token && token <= "Z") || ("0" <= token && token <= "9") {
            // 0 - 9  和 A-Z 范围字符入栈
            postfix.push(token);
        } else if "(" == token  {
            // 遇到开括号，将操作符入栈
            op_stack.push(token);
        } else if ")" == token  {
            // 遇到闭括号，将操作数入栈
            let mut top = op_stack.pop().unwrap();
            while top != "(" {
                postfix.push(top);
                top = op_stack.pop().unwrap();
            }
        } else {
            // 比较符号的优先级来决定操作符号是否加入 postfix
            while (!op_stack.is_empty())
                && (prec[op_stack.peek().unwrap()] >= prec[token]) {
                postfix.push(op_stack.pop().unwrap());
            }
            op_stack.push(token);
        }

        println!("----------t:{:?}",token );
        println!("--op:{:?}",op_stack );
        println!("--out:{:?}\n",postfix );
    }

    // 剩下的操作符入栈
    while !op_stack.is_empty() {
        postfix.push(op_stack.pop().unwrap())
    }

    // 出栈并组成字符串
    let mut postfix_str = "".to_string();
    for c in postfix {
        postfix_str += &c.to_string();
        postfix_str += " ";
    }

    Some(postfix_str)
}

// 中缀转后缀
fn main() {
    let infix = "( A + B ) * ( C + D )";
    let postfix = infix_to_postfix(infix);
    match postfix {
        Some(val) => {
            println!("infix: {infix} -> postfix: {val}");
        },
        None => {
            println!("{infix} is not a corret infix string");
        },
    }
}