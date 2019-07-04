use crate::ListNode;
/// 解题思路：
///
///
#[allow(dead_code)]
impl Solution {
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut head = Box::new(ListNode::new(0));
        let mut ref_head = &mut head;
        let mut carry = 0;

        let mut l1 = l1.as_ref();
        let mut l2 = l2.as_ref();
        while l1.is_some() || l2.is_some() {
            let (x, y) = (l1.map(|n| n.val), l2.map(|n| n.val));
            let (x, y) = (x.unwrap_or(0), y.unwrap_or(0));

            let sum = (x + y + carry) % 10;
            carry = (x + y + carry) / 10;

            ref_head.next = Some(Box::new(ListNode::new(sum)));
            ref_head = ref_head.next.as_mut().unwrap();

            l1 = l1.map(|n| n.next.as_ref()).unwrap_or(None);
            l2 = l2.map(|n| n.next.as_ref()).unwrap_or(None);
        }

        if carry > 0 {
            ref_head.next = Some(Box::new(ListNode::new(carry)));
        }

        head.next
    }
}

pub struct Solution;

#[cfg(test)]
mod test {
    use super::Solution;
    use crate::linkedlist;

    #[test]
    fn test() {
        let l1 = linkedlist![2, 4, 3];
        let l2 = linkedlist![5, 6, 4];
        assert_eq!(Solution::add_two_numbers(l1, l2), linkedlist![7, 0, 8]);

        assert_eq!(
            Solution::add_two_numbers(linkedlist![5], linkedlist![5]),
            linkedlist![0, 1]
        );
    }
}
