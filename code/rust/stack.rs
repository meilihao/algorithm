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

// 仅支持检查`()`
fn par_checker1(par: &str) -> bool {
    // 字符加入 Vec
    let mut char_list = Vec::new();
    for c in par.chars() {
        char_list.push(c);
    }

    let mut index = 0;
    let mut balance = true; // 括号是否匹配(平衡)标示
    let mut stack = Stack::new();
    while index < char_list.len() && balance {
        let c = char_list[index];

        if '(' == c { // 如果为开符号, 入栈
            stack.push(c);
        } else { // 如果为闭符号, 判断栈是否为空
            if stack.is_empty() { // 为空, 所以不匹配
                balance = false;
            } else {
                let _r = stack.pop();
            }
        }

        index += 1;
    }

    // 平衡且栈为空时，括号表达式才是匹配的
    balance && stack.is_empty()
}


// 同时检测多种开闭符号是否匹配
fn par_match(open: char, close: char) -> bool {
    let opens = "([{";
    let closers = ")]}";
    opens.find(open) == closers.find(close) // 位置相同
}

// 支持检查`(), [], {}`
fn par_checker2(par: &str) -> bool {
    let mut char_list = Vec::new();
    for c in par.chars() {
        char_list.push(c);
    }

    let mut index = 0;
    let mut balance = true;
    let mut stack = Stack::new();
    while index < char_list.len() && balance {
        let c = char_list[index];

        // 同时判断三种开符号
        if '(' == c || '[' == c || '{' == c {
            stack.push(c);
        } else {
            if stack.is_empty() {
                balance = false;
            } else {
                // 比较当前括号和栈顶括号是否匹配
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

// 支持检查`(), [], {}`, 允许包含其他字符
fn par_checker3(par: &str) -> bool {
    let mut char_list = Vec::new();
    for c in par.chars() {
        char_list.push(c);
    }

    let mut index = 0;
    let mut balance = true;
    let mut stack = Stack::new();
    while index < char_list.len() && balance {
        let c = char_list[index];

        // 开符号入栈
        if '(' == c || '[' == c || '{' == c {
            stack.push(c);
        }

        // 闭符号则判断是否平衡
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

        // 非括号直接跳过
        index += 1;
    }

    balance && stack.is_empty()
}

fn main() {
    println!("--- stack demo ---");
    let mut s = Stack::new();
    s.push(1); s.push(2); s.push(4);
    println!("top {:?}, size {}",s.peek().unwrap(), s.size());
    println!("pop {:?}, size {}",s.pop().unwrap(), s.size());
    println!("is_empty:{}, stack:{:?}", s.is_empty(), s);

    println!("--- par_checker1 ---");
    // 检查括号是否合法(有开必有闭)
    let sa = "()(())";
    let sb = "()((()";
    let sc = "))";
    let res1 = par_checker1(sa);
    let res2 = par_checker1(sb);
    let res3 = par_checker1(sc);
    println!("sa balanced: {res1}, sb balanced: {res2}, sc balanced: {res3}");

    println!("--- par_checker2 ---");
    let sa = "(){}[]";
    let sb = "(){)[}";
    let res1 = par_checker2(sa);
    let res2 = par_checker2(sb);
    println!("sa balanced: {res1}, sb balanced: {res2}");

    println!("--- par_checker3 ---");
    let sa = "(2+3){func}[abc]";
    let sb = "(2+3)*(3-1";
    let res1 = par_checker3(sa);
    let res2 = par_checker3(sb);
    println!("sa balanced: {res1}, sb balanced: {res2}");
}