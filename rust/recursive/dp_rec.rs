// https://github.com/QMHTMY/RustBook/blob/main/code/chapter04/dp_rec.rs
// 斐波那契数列: 递归和动态规划都可解
fn fibnacci_rec(n: u32) -> u32 {
    if n == 1 || n == 2 {
        return 1;
    } else {
        fibnacci_rec(n - 1) + fibnacci_rec(n - 2)
    }
}

fn fibnacci_dp(n: u32) -> u32 {
    // 只用两个位置来保存值，节约内存
    let mut dp = [1, 1];

    for i in 2..=n {
        let idx1 = (i % 2) as usize; // 保存当前结果的索引
        let idx2 = ((i - 1) % 2) as usize; // 前1的索引
        let idx3 = ((i - 2) % 2) as usize; // 前2的索引
        // println!("idx1,idx2,idx3: {},{},{}", idx1,idx2,idx3);
        dp[idx1] = dp[idx2] + dp[idx3];
    }

    dp[((n-1) % 2) as usize] // n-1: 上面for结束后其实n已加1
}

fn main() {
    println!("fib(10): {}", fibnacci_rec(10));
    println!("fib(10): {}", fibnacci_dp(10));
}