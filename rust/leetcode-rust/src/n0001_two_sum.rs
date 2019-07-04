use std::collections;

#[allow(dead_code)]
impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut h = collections::HashMap::new();
        let mut res = vec![];
        for (index, num) in nums.iter().enumerate() {
            let key = target - *num;
            if h.contains_key(&key) {
                let another = h.get(&key).unwrap();
                res = vec![*another as i32, index as i32];
            }
            h.insert(*num, index);
        }
        res
    }
}

pub struct Solution;

#[cfg(test)]
mod test {
    use super::Solution;
    #[test]
    fn test() {
        let nums = vec![2, 7, 11, 15];
        let target = 9;
        let res = Solution::two_sum(nums, target);
        assert_eq!(res, vec![0, 1])
    }
}
