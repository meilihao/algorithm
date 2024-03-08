// https://github.com/QMHTMY/RustBook/blob/main/code/chapter04/nums_sum34.rs
// nums_sum34.rs
// 所有参数都放到递归函数里，且函数的最后一行一定是递归，所以又叫尾递归
fn nums_sum3(sum: i32, nums: &[i32]) -> i32 {
    if 1 == nums.len() {
        sum + nums[0]
    } else {
        // 使用 sum 来接收中间计算值
        nums_sum3(sum + nums[0], &nums[1..])
    }
}

fn nums_sum4(sum: i32, nums: &[i32]) -> i32 {
    if 1 == nums.len() {
        sum + nums[0]
    } else {
        nums_sum4(sum + nums[nums.len() - 1], &nums[..nums.len() - 1])
    }
}
fn main() {
    let nums = [2,1,7,4,5];
    let sum1 = nums_sum3(0, &nums);
    let sum2 = nums_sum4(0, &nums);
    println!("sum1 is {sum1}, sum2 is {sum2}");
}