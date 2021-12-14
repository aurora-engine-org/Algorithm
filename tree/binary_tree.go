package tree

import (
	"fmt"
	"math"
)

// BinarySearchTree 创建排序二叉树
func BinarySearchTree(root *node, child *node) {
	if child.value < root.value {
		if root.left == nil {
			root.left = child
			return
		}
		BinarySearchTree(root.left, child)
		return
	}
	if child.value > root.value {
		if root.right == nil {
			root.right = child
			return
		}
		BinarySearchTree(root.right, child)
		return
	}

	if child.value == root.value {
		return
	}
}

// AvlTree 创建平衡二叉树
func AvlTree(root *node, child *node) *node {
	if root == nil {
		return nil
	}
	if child.value > root.value {
		if root.right == nil {
			child.parent = root //记住自己的父节点
			root.right = child
			return root
		}
		g := AvlTree(root.right, child)
		//计算深度
		if b, h := isAvl(g); !b {
			if h > 0 {
				//大于0 左子树过高
				fmt.Println("root right 左子树过高 当前根节点value:", g.value)
				//判断当前根节点 发生不平衡的节点在左子树的那个节点上，检测过程最多走两步

			}
			if h < 0 {
				//小于0 右子树过高
				fmt.Println("root right 右子树过高 当前根节点value:", g.value)
			}
		}
	}
	if child.value < root.value {
		if root.left == nil {
			child.parent = root //记住自己的父节点
			root.left = child
			return root
		}
		g := AvlTree(root.left, child)
		//计算深度
		if b, h := isAvl(g); !b {
			if h > 0 {
				//大于0 左子树过高
				fmt.Println("root left 左子树过高 当前根节点value:", g.value)

				//拿到 根节点的左孩子
				gl := g.left

				//拿到根节点的父节点
				gp := g.parent

				//判断是 左左类型 还是 左右类型
				if gl.left != nil && gl.right == nil {
					gl.right = g
					g.left = nil
					g.parent = gl
					gp.left = gl
					gl.parent = gp
				}

				if gl.left == nil && gl.right != nil {
					gp.left = gl.right
					gl.right.parent = gp

					gl.right.right = g
					g.parent = gl.right
					g.left = nil

					gl.right.left = gl
					gl.parent = gl.right
					gl.right = nil
				}

			}
			if h < 0 {
				//小于0 右子树过高
				fmt.Println("root left 右子树过高 当前根节点value:", g.value)
			}
		}
	}
	return root
}

// 检测是否为avl树，并且返回高度差
func isAvl(root *node) (bool, int) {
	h := 0  //当前节点层级深度
	ln := 0 //左子树深度
	rn := 0 //右子树深度
	h++     //计算当前一步
	if root.left != nil {
		ln = AvlDepth(root.left)
	}
	if root.right != nil {
		rn = AvlDepth(root.right)
	}
	//检查当前节点的平衡性
	v := math.Abs(float64(ln - rn))
	if v == 0 || v == 1 {
		return true, int(v)
	}
	return false, ln - rn //以左子树 减 右子树为标准来区分  大于0 左子树过高，小于0 右子树过高
}

// BinaryTreeInorderTraversal 二叉树中序遍历 先根
func BinaryTreeInorderTraversal(root *node) {
	if root == nil {
		return
	}
	BinaryTreeInorderTraversal(root.left)
	fmt.Println(root.value)
	BinaryTreeInorderTraversal(root.right)
}

// MaxDepth 计算二叉树深度
func MaxDepth(root *node) int {
	h := 0  //当前节点层级深度
	ln := 0 //左子树深度
	rn := 0 //右子树深度
	h++     //计算当前一步
	if root.left != nil {
		ln = MaxDepth(root.left)
	}
	if root.right != nil {
		rn = MaxDepth(root.right)
	}
	if ln > rn { //对比那个子树深度最大就返回那个
		return ln + h
	} else {
		return rn + h
	}
}

// AvlDepth 平衡二叉树计算深度
func AvlDepth(root *node) int {
	h := 0  //当前节点层级深度
	ln := 0 //左子树深度
	rn := 0 //右子树深度
	h++     //计算当前一步
	if root.left != nil {
		ln = AvlDepth(root.left)
	}
	if root.right != nil {
		rn = AvlDepth(root.right)
	}
	if ln > rn { //对比那个子树深度最大就返回那个
		return ln + h
	} else {
		return rn + h
	}
}

// Judge 递归判断不平衡节点所在位置
func Judge(root *node) *node {
	if root.left != nil && root.right != nil {
		Judge(root.left)
		Judge(root.right)
	}
	if root.left != nil && root.right == nil {
		Judge(root.left)
	}
	if root.left == nil && root.right != nil {
		Judge(root.right)
	}
	return root
}
