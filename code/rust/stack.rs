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
