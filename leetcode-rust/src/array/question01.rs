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
    if strs.len() == 0 {
        return String::new();
    }

    let min_len = strs.iter().map(|s| s.len()).min();
    if let Some(0) = min_len {
        return String::new();
    }

    let min_len = min_len.unwrap();
    let mut count = 0;
    let first = strs[0].clone();
    for i in 0..min_len {
        let tmp_size = strs
            .iter()
            .map(|s| s.as_bytes())
            .map(|arr| arr[i])
            .collect::<HashSet<_>>()
            .len();
        if tmp_size != 1 {
            if count == 0 {
                return String::new();
            } else {
                let tmp = first.get(0..count).unwrap();
                return String::from(tmp);
            }
        }
        count += 1;
    }
    if count == 0 {
        return String::new();
    } else {
        let tmp = first.get(0..count).unwrap();
        return String::from(tmp);
    }
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
}
