package main

import (
    "fmt"
)

type SplayNode struct {
    key    int
    left   *SplayNode
    right  *SplayNode
    parent *SplayNode
}

type SplayTree struct {
    root *SplayNode
}

func insert(tree *SplayTree, node *SplayNode, key int) {
    if node.key == key {
        return
    }
    if key < node.key {
        if node.left == nil {
            node.left = &SplayNode{key: key, parent: node}
            x := splay(node.left)
            tree.root = x
            tree.root.parent = nil
            return
        }
        insert(tree, node.left, key)
    } else {
        if node.right == nil {
            node.right = &SplayNode{key: key, parent: node}
            x := splay(node.right)
            tree.root = x
            tree.root.parent = nil
            return
        }
        insert(tree, node.right, key)
    }
}

func Insert(tree *SplayTree, key int) {
    if tree.root == nil {
        tree.root = &SplayNode{key: key}
        return
    }
    insert(tree, tree.root, key)
}

func Search(tree *SplayTree, key int) *SplayNode {
    if tree.root == nil {
        return nil
    }
    node := search(tree, tree.root, key)
    tree.root = splay(node)
    return node
}

func search(tree *SplayTree, node *SplayNode, key int) *SplayNode {
    if node == nil {
        return node
    }

    if key == node.key {
        tree.root = splay(node)
        return node
    }
    if key < node.key {
        n := search(tree, node.left, key)
        tree.root = splay(n)
        return n
    }
    n := search(tree, node.left, key)
        tree.root = splay(n)
        return n
}

func Delete(tree *SplayTree, key int) {
    if tree.root == nil {
        return
    }
    node := search(tree, tree.root, key)
    if node == nil {
        return
    }
    tree.root = splay(node)
    if node.left == nil {
        tree.root = node.right
        return
    }
    x := node.left
    x.parent = nil
    for x.right != nil {
        x = x.right
    }
    x = splay(x)
    x.right = node.right
    tree.root = x
}

func isZig(node *SplayNode) bool {
    return node.parent.parent == nil
}

func isZigZig(node *SplayNode) bool {
    return (node.parent.left == node && node.parent.parent.left == node.parent) ||
        (node.parent.right == node && node.parent.parent.right == node.parent)
}

func isZigZag(node *SplayNode) bool {
    return (node.parent.right == node && node.parent.parent.left == node.parent) ||
        (node.parent.left == node && node.parent.parent.right == node.parent)
}

func splay(node *SplayNode) *SplayNode {
    if node == nil {
        panic("Node is nil")
    }
    for node.parent != nil {
        if isZig(node) {
            zig(node)
            break
        } else if isZigZig(node) {
            zigZig(node)
        } else if isZigZag(node) {
            zigZag(node)
        }
    }
    return node
}

func zigZag(node *SplayNode) {
    if node == nil {
        panic("Node is nil")
    }
    parent := node.parent

    if parent.right == node {
        rotateLeft(parent)
        rotateRight(parent)
    }
    if parent.left == node {
        rotateRight(parent)
        rotateLeft(parent)
    }
}

func zigZig(node *SplayNode) {
    if node == nil {
        panic("Node is nil")
    }
    parent := node.parent
    grandParent := parent.parent

    if parent.left == node && grandParent.left == parent {
        rotateRight(grandParent)
        rotateRight(parent)
    }
    if parent.right == node && grandParent.right == parent {
        rotateLeft(grandParent)
        rotateLeft(parent)
    }
}

func zig(node *SplayNode) *SplayNode {
    if node.parent.left == node {
        rotateRight(node.parent)
    }
    if node.parent.right == node {
        rotateLeft(node.parent)
    } else {
        panic("Node is not a child of its parent")
    }
    return node
}

func rotateRight(node *SplayNode) {
    parent := node.parent
    x := node.left
    if x == nil {
        return
    }
    node.left = x.right
    if x.right != nil {
        x.right.parent = node
    }
    x.parent = parent
    if parent != nil {
        if node == parent.left {
            parent.left = x
        } else {
            parent.right = x
        }
    }
    x.right = node
    node.parent = x
}

func rotateLeft(node *SplayNode) {
    parent := node.parent
    x := node.right
    if x == nil {
        return
    }
    node.right = x.left
    if x.left != nil {
        x.left.parent = node
    }
    x.parent = parent
    if parent != nil {
        if node == parent.left {
            parent.left = x
        } else {
            parent.right = x
        }
    }
    x.left = node
    node.parent = x
}

func printLevels(root *SplayNode) {
    if root == nil {
        return
    }

    queue := []*SplayNode{root}

    for len(queue) > 0 {
        levelSize := len(queue)

        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]

            fmt.Print(node.key, " ")

            if node.left != nil {
                queue = append(queue, node.left)
            }
            if node.right != nil {
                queue = append(queue, node.right)
            }
        }

        fmt.Println()
    }
}

func inOrder(node *SplayNode) {
    if node == nil {
        return
    }
    inOrder(node.left)
    fmt.Println(node.key)
    inOrder(node.right)
}
func main() {
    tree := SplayTree{}
    Insert(&tree, 10)
    Insert(&tree, 20)
    Insert(&tree, 30)
    Insert(&tree, 15)
    Insert(&tree, 25)
    printLevels(tree.root)
    println("Inorder Traversal")
    inOrder(tree.root)
    println()
    println(Search(&tree, 15).key)
    println()
    inOrder(tree.root)

    println()
    Delete(&tree, 15)
    println()
    inOrder(tree.root)

}
