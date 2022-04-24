// 双端队列
#[derive(Debug)]
struct Deque<T> {
    cap: usize,   // 容量
    data: Vec<T>, // 数据容器
}

impl<T> Deque<T> {
    fn new(cap: usize) -> Self {
        Deque {
            cap: cap,
            data: Vec::with_capacity(cap),
        }
    }

    // Vec 末尾为队首
    fn add_front(&mut self, val: T) -> Result<(), String> {
        if self.size() == self.cap {
            return Err("No space avaliable".to_string());
        }
        self.data.push(val);

        Ok(())
    }

    // Vec 首部为队尾
    fn add_rear(&mut self, val: T) -> Result<(), String> {
        if self.size() == self.cap {
            return Err("No space avaliabl".to_string());
        }
        self.data.insert(0, val);

        Ok(())
    }

    // 从队首移除数据
    fn remove_front(&mut self) -> Option<T> {
        if self.size() > 0 {
            self.data.pop()
        } else {
            None
        }
    }

    // 从队尾移除数据
    fn remove_rear(&mut self) -> Option<T> {
        if self.size() > 0 {
            Some(self.data.remove(0))
        } else {
            None
        }
    }

    fn is_empty(&self) -> bool {
        // self.size() == 0
        self.data.len() == 0
    }

    fn size(&self) -> usize {
        self.data.len()
    }
}

// 游戏: 烫手山芋, 类似约瑟夫问题
// 初始山芋在对首, 向后移动, 每移动n次, 删除一次对首
fn hot_potato(names: Vec<&str>, num: usize) -> &str {
    // 初始化队列、名字入队
    let mut q = Deque::new(names.len());
    for name in names {
        let _nm = q.add_rear(name); // 从尾部入队
    }

    println!("deque:{:?}", q);

    while q.size() > 1 {
        // 出入栈名字，相当于传递山芋
        for _i in 0..num {
            let name = q.remove_front().unwrap(); // 首部出队
            let _rm = q.add_rear(name);
        }

        // 出入栈达到 num 次，删除一人
        let _rm = q.remove_front();
    }

    q.remove_front().unwrap() // 剩下的人
}

fn main() {
    let mut d = Deque::new(4);
    let _r1 = d.add_front(1); let _r2 = d.add_front(2);
    let _r1 = d.add_rear(3); let _r2 = d.add_rear(4);
    if let Err(error) = d.add_front(5) {
        println!("add_front error: {error}");
    }

    if let Some(data) = d.remove_rear() {
        println!("remove data {data}");
    } else {
        println!("empty deque");
    }

    println!("size: {}, is_empyt: {}", d.size(), d.is_empty());
    println!("content: {:?}", d);

    let names = vec!["Shieber", "Tom", "Kew", "Lisa", "Marry", "Bob"];
    let rem = hot_potato(names, 2);
    println!("The left person is {rem}");
}