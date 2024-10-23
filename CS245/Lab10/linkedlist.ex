@moduledoc "
Define a structure to implement the node of a linked list (the structure should have a data payload and a pointer to the next node)
"
defmodule MyNode do
  defstruct data: nil, next: nil
end

@moduledoc "
Write, in Elixir, code to create and provide functionality for you own linked list implementation. You should write code to do the following:
Define a structure to hold a linked list. The contents of the structure should be a pointer to the head of the linked list and the number of nodes in the linked list.
Create a method to add a new node at the front of a linked list. The method show take two arguments: a data item and a linked list struct. The returned value should be a linked list struct.
Create a method to print the contents of one of your linked lists
Create a method that adds a data item to the tail of a linked list. It should take a data item and a linked list struct as arguments. It should return a linked list struct.
Create a function that appends one of your lists to another of your lists. The method should take two list structs as arguments and return a list struct.
"
defmodule LinkedList do
  defstruct head: nil, length: 0

  def add_front(data, list) do
    new_node = %MyNode{data: data, next: list.head}
    %LinkedList{head: new_node, length: list.length + 1}
  end

  def add_tail(data, list) do
    new_node = %MyNode{data: data, next: nil}
    if list.head == nil do
      %LinkedList{head: new_node, length: list.length + 1}
    else
      add_tail_helper(new_node, list.head)
      %LinkedList{head: list.head, length: list.length + 1}
    end
  end

  defp add_tail_helper(new_node, node) do
    if node.next == nil do
      node.next = new_node
    else
      add_tail_helper(new_node, node.next)
    end
  end

  def append(list1, list2) do
    if list1.head == nil do
      list2
    else
      append_helper(list1.head, list2.head)
      %LinkedList{head: list1.head, length: list1.length + list2.length}
    end
  end

  defp append_helper(node1, node2) do
    if node1.next == nil do
      node1.next = node2
    else
      append_helper(node1.next, node2)
    end
  end

  def print_list(list) do
    if list.head == nil do
      IO.puts("Empty list")
    else
      print_list_helper(list.head)
    end
  end

  defp print_list_helper(node) do
    IO.puts(node.data)
    if node.next != nil do
      print_list_helper(node.next)
    end
  end
end

a = LinkedList.add_front("a", nil)
b = LinkedList.add_front("b", a)
c = LinkedList.add_tail("c", b)
LinkedList.print_list(b)
LinkedList.print_list(c)
LinkedList.append(c, b)
