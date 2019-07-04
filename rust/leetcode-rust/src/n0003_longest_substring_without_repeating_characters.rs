use std::collections::HashMap;

#[allow(dead_code)]
impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        let mut h = HashMap::new();
        let mut flag = 0;
        let mut ans = 0;
        for (i, c) in s.chars().enumerate() {
            if let Some(t) = h.get(&c) {
                if *t > flag {
                    flag = *t;
                }
            }
            if ans < i - flag + 1 {
                ans = i - flag + 1
            }
            h.insert(c, i + 1);
        }
        ans as i32
    }
}

pub struct Solution;

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test() {
        let s = String::from("abcda");
        let expect = 4;
        let res = Solution::length_of_longest_substring(s);
        assert_eq!(res, expect);

        let s = String::from("abba");
        let expect = 2;
        let res = Solution::length_of_longest_substring(s);
        assert_eq!(res, expect);

        let s = String::from("");
        let expect = 0;
        let res = Solution::length_of_longest_substring(s);
        assert_eq!(res, expect);
    }
}
