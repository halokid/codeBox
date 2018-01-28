#coding=utf-8

import feedparser
import re

def get_word_counts(url):
  """
  :param url:   feed url
  :return: list
  """
  d = feedparser.parse(url)
  wc = {}

  for e in d.entries:
    if 'summary' in e:
      summary = e.summary
    else:
      summary = e.description

    words = getwords(e.title + ' ' + summary)
    for word in words:
      wc.setdefault(word, 0)
      wc[word] += 1

  return d.feed.title, wc



def getwords(html):
  txt = re.compile(r'<[^>] +>').sub('', html)
  words = re.compile(r'[^A-Z^a-z]+').split(txt)

  return [word.lower() for word in words if word != '']





apcount = {}
wordcounts = {}
feedlist = [line for line in file('feedlist.txt')]
for feedurl in feedlist:
  title, wc = get_word_counts(feedurl)
  wordcounts[title] = wc
  for word, count in wc.items():
    apcount.setdefault(word, 0)   #是计算包含这些单词的博客数目
    if count > 1:
      apcount[word] += 1




