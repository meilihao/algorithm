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

// 将10进制转为二进制
fn divide_by_two(mut dec_num: u32) -> String {
    // 用栈来保存余数 rem
    let mut rem_stack = Stack::new();

    // 余数 rem 入栈
    while dec_num > 0 {
        let rem = dec_num % 2;
        rem_stack.push(rem);
        dec_num /= 2;
    }

    // 栈中元素出栈组成字符串
    let mut bin_str = "".to_string();
    while !rem_stack.is_empty() {
        let rem = rem_stack.pop().unwrap().to_string();
        bin_str += &rem;
    }

    bin_str
}

// 支持其他进制, 最高是16进制
fn base_converter(mut dec_num: u32, base: u32) -> String {
    // digits 对应各种余数的字符形式，尤其是 10 - 15
    let digits = ['0', '1', '2', '3', '4', '5', '6', '7',
                 '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'];

    // 余数入栈
    let mut rem_stack = Stack::new();
    while dec_num > 0 {
        let rem = dec_num % base;
        rem_stack.push(rem);
        dec_num /= base;
    }

    // 余数出栈并取对应字符来拼接成字符串
    let mut base_str = "".to_string();
    while !rem_stack.is_empty() {
        let rem = rem_stack.pop().unwrap() as usize;
        base_str += &digits[rem].to_string();
    }

    base_str
}

fn main() {
    let bin_str: String = divide_by_two(10);
    println!("10 is b{bin_str}");

    let bin_str: String = base_converter(10, 2);
    let hex_str: String = base_converter(43, 16);
    println!("10 is b{bin_str}, 43 is 0x{hex_str}");
}