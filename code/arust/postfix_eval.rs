use std::collections::HashMap;

#[derive(Debug)]
struct Stack<T> {
    size: usize,  // 栈中元素个数
    data: Vec<T>, // 栈数据
}

impl<T> Stack<T> {
    fn new() -> Self { // 初始化空栈
        Stack {
            size: 0,
            data: Vec::new()
        }
    }

    fn push(&mut self, val: T) {
        self.data.push(val); // 数据保存在 Vec 末尾
        self.size += 1;
    }

    fn pop(&mut self) -> Option<T> {
        if self.size == 0 { return None; }
        self.size -= 1; // size减 1 后再弹出数据
        self.data.pop()
    }

    fn peek(&self) -> Option<&T> { // 数据不能移动，只能返回引用
        if self.size == 0 { return None; }
        self.data.get(self.size - 1)
    }

    fn is_empty(&self) -> bool {
        0 == self.size
    }

    fn size(&self) -> usize {
        self.size
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
    // 使用了一个名为 prec 的 HashMap 来保存操作符优先级，它将每个运算符映射为一个整数，用于与其他运算符优先级进行比较。左括号赋予最低的优先级，这样，与其进行比较的任何运算符都具有更高的优先级
    let mut prec = HashMap::new();
    prec.insert("(", 1); prec.insert(")", 1);
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

fn postfix_eval(postfix: &str) -> Option<i32> {
    // 少于五个字符，不是有效的后缀表达式，因为表达式
    // 至少两个操作数加一个操作符，还需要两个空格隔开
    if postfix.len() < 5 { return None; }

    let mut op_stack = Stack::new(); // 操作数栈
    for token in postfix.split_whitespace() {
        if "0" <= token  && token <= "9" {
            op_stack.push(token.parse::<i32>().unwrap());
        } else {
            // 对于减法和除法，顺序有要求
            // 所以先出栈的是第二个操作数
            let op2 = op_stack.pop().unwrap();
            let op1 = op_stack.pop().unwrap();
            let res = do_calc(token, op1, op2);
            op_stack.push(res);
        }
    }

    Some(op_stack.pop().unwrap())
}

// 执行数学运算
fn do_calc(op: &str, op1: i32, op2: i32) -> i32 {
    if "+" == op  {
        op1 + op2
    } else if "-" == op {
        op1 - op2
    } else if "*" == op {
        op1 * op2
    } else {
        if 0 == op2 {
            panic!("ZeroDivisionError: Invalid operation!");
        }
        op1 / op2
    }
}

fn main() {
    let infix = "( 1 + 2 ) * ( 1 + 2 )";
    let postfix = infix_to_postfix(infix);
    match postfix {
        Some(ref val) => {
            println!("infix: {infix} -> postfix: {val}");
        },
        None => {
            println!("{infix} is not a corret infix string");
        },
    }

    let postfix = postfix.unwrap();
    let res = postfix_eval(&postfix);
    match res {
        Some(val) => println!("res is {val}"),
        None => println!("{postfix} isn't a corret postfix"),
    }
}