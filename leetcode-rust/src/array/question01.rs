use std::collections::{HashMap, HashSet};

fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut map = HashMap::new();
    for (i, v) in nums.iter().enumerate() {
        let tmp = target - v;
        let other = map.get(&tmp);

        if let Some(n) = other {
            return vec![i as i32, *n as i32];
        } else {
            map.insert(v, i);
        }
    }
    vec![]
}

fn longest_common_prefix(strs: Vec<String>) -> String {
    if strs.is_empty() {
        return String::new();
    }

    let first = &strs[0];
    let min_len = strs.iter().map(|s| s.len()).min().unwrap_or(0);

    for i in 0..min_len {
        let byte = first.as_bytes()[i];
        if !strs[1..].iter().all(|s| s.as_bytes()[i] == byte) {
            return first[..i].to_string();
        }
    }

    first[..min_len].to_string()
}

fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
    if nums.is_empty() {
        return 0;
    }

    let mut slow = 1;
    for fast in 1..nums.len() {
        if nums[fast] != nums[fast - 1] {
            nums[slow] = nums[fast];
            slow += 1;
        }
    }
    return slow as i32;
}

// 居然内存占用排在末尾
// https://leetcode.cn/problems/search-insert-position
fn search_insert(nums: Vec<i32>, target: i32) -> i32 {
    let mut left = 0;
    let mut right = nums.len();
    let mut mid = (left + right) / 2;
    let mut idx = right;
    while left < right {
        if target > nums[mid] {
            left = mid + 1;
        } else {
            idx = std::cmp::min(idx, mid);
            right = mid;
        }
        mid = (left + right) / 2;
    }
    return idx as i32;
}

#[cfg(test)]
mod tests {
    use super::*;
    use rstest::rstest;

    #[rstest]
    #[case(vec![2,7,11,15], 9,vec![0,1])]
    #[case(vec![3,2,4], 6,vec![1,2])]
    fn test_two_sum(#[case] input: Vec<i32>, #[case] target: i32, #[case] want: Vec<i32>) {
        let mut ans = two_sum(input, target);
        ans.sort();
        assert_eq!(want, ans)
    }

    #[rstest]
    #[case(vec!["flower","flow","flight"], "fl".to_string())]
    #[case(vec!["dog","racecar","car"], "".to_string())]
    #[case(vec!["a"], "a".to_string())]
    fn test_longest_common_prefix(#[case] input: Vec<&str>, #[case] want: String) {
        let strs = input.into_iter().map(|s| s.to_string()).collect();
        let ans = longest_common_prefix(strs);
        assert_eq!(want, ans)
    }

    #[rstest]
    #[case(vec![1,3,5,6], 5,2)]
    #[case(vec![1,3,5,6], 2,1)]
    fn test_search_insert(#[case] input: Vec<i32>, #[case] target: i32, #[case] want: i32) {
        let ans = search_insert(input, target);
        assert_eq!(want, ans)
    }
}
