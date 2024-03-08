// https://github.com/QMHTMY/RustBook/blob/main/code/chapter04/move2tower.rs
// p : pole 杆
// height: 初始盘数

// 逻辑:
// 1. 借助目标杆将height-1个盘子移动到中间杆
// 2. 将最后一个盘子移动到目标杆
// 3. 借助起始杆将height-1个盘子从中间杆移动到目标杆
fn move2tower(height: u32, src_p: &str, des_p: &str, mid_p: &str) {
    if height >= 1 {
        move2tower(height - 1, src_p, mid_p, des_p);
        println!("moving disk[{height}] from {src_p} to {des_p}");
        move2tower(height - 1, mid_p, des_p, src_p);
    }
}

fn main() {
    move2tower(1, "A", "B", "C");
    move2tower(2, "A", "B", "C");
    move2tower(3, "A", "B", "C");
    move2tower(4, "A", "B", "C");
}