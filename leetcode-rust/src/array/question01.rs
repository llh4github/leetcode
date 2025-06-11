use std::collections::HashMap;

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
}
