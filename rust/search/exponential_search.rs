// https://github.com/QMHTMY/RustBook/blob/main/code/chapter05/exponential_search.rs
// exponential_search.rs 指数查找是二分查找的变形. 它划分中值的方法不是使用平均或插值而是用指数函数来估计，这样可以快速找到上界，加快查找，该算法适合已排序且无边界的数据

fn binary_search(nums: &[i32], num: i32) -> bool {
    let mut low = 0;
    let mut high = nums.len() - 1;
    let mut found = false;

    // 注意是 <= 不是 <
    while low <= high && !found {
        let mid: usize = (low + high) >> 1;
        if num == nums[mid] {
            found = true;
        } else if num < nums[mid] {
            high = mid - 1;
        } else {
            low = mid + 1;
        }
    }

    found
}

fn exponential_search(nums: &[i32], target: i32) -> bool {
    let size = nums.len();
    if size == 0 {
        return false;
    }

    // 逐步找到上界
    let mut high = 1usize;
    while high < size && nums[high] < target { // 在不越界的情况下评估high的范围
        high <<= 1;
    }

    //  上界的一半一定可以作为下界  // 因为按2的指数定位上界, 上面while结束后`nums[high/2]<target<= nums[high]或已越界`
    let low = high >> 1;

    binary_search(&nums[low..size.min(high+1)], target) // high+1即包含nums[high]
}

fn main() {
    let nums = [1,9,10,15,16,17,19,23,27,28,29,30,32,35];
    let target = 27;
    let found = exponential_search(&nums, target);
    println!("{target} in nums: {found}");
}