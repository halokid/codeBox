#coding=utf-8

import feedparser
import re


def debug(ist):
  print ist

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



#返回所有匹配的 word 的一个列表
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

debug(wordcounts)
print "--------------------------------------------------\n"
debug(apcount)


#某个word 在所有的 feedlist 里面出现的概率比
wordlist = []
for w, bc in apcount.items():
  frac = float(bc) / len(feedlist)
  print frac
  if frac > 0.1 and frac < 0.5:
    wordlist.append(w)

debug(wordlist)


out = file('blogdata.txt', 'w')
out.write('Blog')
for word in wordlist:
  out.write('\t%s' % word)
out.write('\n')

for blog, wc in wordcounts.items():
  out.write(blog)
  for word in wordlist:
    if word in wc: out.write('\t%d' % wc[word])
    else: out.write('\t0')
  out.write('\n')


















