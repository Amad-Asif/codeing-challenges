package examples

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeString(t *testing.T) {
	assert.Equal(t, DecodeString("3[a2[c]]"), "accaccacc")
	assert.Equal(t, DecodeString("3[a]2[bc]"), "aaabcbc")
	assert.Equal(t, DecodeString("2[abc]3[cd]ef"), "abcabccdcdcdef")
	assert.Equal(t, DecodeString("11[ab]ef"), "abababababababababababef")
	//fmt.Println(examples.DecodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef") == "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef")
}

func TestCompressString(t *testing.T) {
	s1, l1 := CompressString([]byte("aabbccc"))
	assert.Equal(t, s1, "a2b2c3")
	assert.Equal(t, l1, 6)

	s2, l2 := CompressString([]byte("abbbbbbbbbbbb"))
	assert.Equal(t, s2, "ab12")
	assert.Equal(t, l2, 4)

	s3, l3 := CompressString([]byte("a"))
	assert.Equal(t, s3, "a")
	assert.Equal(t, l3, 1)

}

func TestAddTwoNumbers(t *testing.T) {
	l1 := ListNode{
		Val: 9, Next: &ListNode{
			Val: 9, Next: &ListNode{
				Val: 9, Next: &ListNode{
					Val: 9, Next: &ListNode{
						Val: 9, Next: &ListNode{
							Val: 9, Next: &ListNode{
								Val: 9, Next: nil}}}}}}}
	l2 := ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}
	_ = AddTwoNumbers(&l1, &l2)

	l3 := ListNode{Val: 9, Next: &ListNode{9, &ListNode{1, nil}}}
	l4 := ListNode{Val: 1}
	_ = AddTwoNumbers(&l3, &l4)
}

func TestGetLongestSubstringNoRepetitions(t *testing.T) {
	res0 := GetLongestSubstringNoRepetitions("aab")
	res1 := GetLongestSubstringNoRepetitions("abcabcbb")
	res2 := GetLongestSubstringNoRepetitions("pwwkew")
	res3 := GetLongestSubstringNoRepetitions("bbbbb")
	assert.Equal(t, res1, 3)
	assert.Equal(t, res2, 3)
	assert.Equal(t, res3, 1)
	assert.Equal(t, res0, 2)
}
