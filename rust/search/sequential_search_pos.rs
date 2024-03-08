// https://github.com/QMHTMY/RustBook/blob/main/code/chapter05/sequential_search_pos.rs
// sequential_search_pos.rs

fn sequential_search_pos(nums: &[i32], num: i32) -> Option<usize> {
    let mut pos: usize = 0;
    let mut found = false;

    while pos < nums.len() && !found {
        if num == nums[pos] {
            found = true; // 没必要break, 因为while里的`!found`
        } else {
            pos += 1;
        }
    }

    if found {
        Some(pos)
    } else {
        None
    }
}

fn main() {
    let num = 8;
    let nums = [9,3,7,4,1,6,2,8,5];
    match sequential_search_pos(&nums, num) {
        Some(pos) => println!("the index of {num} is {pos}"),
        None => println!("{num} is not in nums"),
    }
}