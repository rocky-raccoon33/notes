---
title: LRU cache
comments: true
description: Least Recently Used (LRU) cache.
date: 2023-03-12
category: 数据结构
section: blog
tags:
    - 数据结构
---

## LRU Cache implemented in C++

```cpp title="LRLCache.cpp" linenums="1"
#include <iostream>
#include <list>
#include <unordered_map>

using namespace std;

class LRUCache
{
	size_t m_capacity;
	unordered_map<int, list<pair<int, int>>::iterator> m_map;
	list<pair<int, int>> m_list;

public:
	LRUCache(size_t capacity) : m_capacity(capacity) {}

	int get(int key)
	{
		auto found_iter = m_map.find(key);
		if (found_iter == m_map.end())
		{
			return -1;
		}
		m_list.splice(m_list.begin(), m_list, found_iter->second);
		return found_iter->second->second;
	}
	void set(int key, int val)
	{
		auto found_iter = m_map.find(key);
		if (found_iter != m_map.end())
		{
			m_list.splice(m_list.begin(), m_list, found_iter->second);
			found_iter->second->second = val;
			return;
		}
		if (m_map.size() == m_capacity)
		{
			int key_to_del = m_list.back().first;
			m_list.pop_back();
			m_map.erase(key_to_del);
		}
		m_list.emplace_front(key, val);
		m_map[key] = m_list.begin();
	}
};
```