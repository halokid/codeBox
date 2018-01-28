#coding=utf-8

critics={'Lisa Rose': {'Lady in the Water': 2.5, 'Snakes on a Plane': 3.5,
                       'Just My Luck': 3.0, 'Superman Returns': 3.5, 'You, Me and Dupree': 2.5,
                       'The Night Listener': 3.0},
         'Gene Seymour': {'Lady in the Water': 3.0, 'Snakes on a Plane': 3.5,
                          'Just My Luck': 1.5, 'Superman Returns': 5.0, 'The Night Listener': 3.0,
                          'You, Me and Dupree': 3.5},
         'Michael Phillips': {'Lady in the Water': 2.5, 'Snakes on a Plane': 3.0,
                              'Superman Returns': 3.5, 'The Night Listener': 4.0},
         'Claudia Puig': {'Snakes on a Plane': 3.5, 'Just My Luck': 3.0,
                          'The Night Listener': 4.5, 'Superman Returns': 4.0,
                          'You, Me and Dupree': 2.5},
         'Mick LaSalle': {'Lady in the Water': 3.0, 'Snakes on a Plane': 4.0,
                          'Just My Luck': 2.0, 'Superman Returns': 3.0, 'The Night Listener': 3.0,
                          'You, Me and Dupree': 2.0},
         'Jack Matthews': {'Lady in the Water': 3.0, 'Snakes on a Plane': 4.0,
                           'The Night Listener': 3.0, 'Superman Returns': 5.0, 'You, Me and Dupree': 3.5},
         'Toby': {'Snakes on a Plane':4.5,'You, Me and Dupree':1.0,'Superman Returns':4.0}}

from math import sqrt

#欧几里德距离
def sim_distance(prefs, person1, person2):
  si = {}
  for item in prefs[person1]:
    if item in prefs[person2]:
      si[item] = 1

  if len(si) == 0:
    return 0

  sum_of_squares=sum([pow(prefs[person1][item]-prefs[person2][item],2)
                      for item in prefs[person1] if item in prefs[person2]])

  return 1 / (1 + sqrt(sum_of_squares))



#皮尔逊相关度


sim_person = None
#推荐物品
def getRecommendations(prefs, person, similarity=sim_person):
  """
  获得最相似的人评价最高的物品
  :param prefs:
  :param person:
  :param similarity:
  :return:
  """
  totals = {}
  sim_sums = {}

  for other in prefs:
    if other == person: continue
    sim = similarity(prefs, person, other)

    if sim <= 0: continue
    for item in prefs[other]:
      #指定的person还没有看过的判断
      if item not in prefs[person] or prefs[person][item] == 0:
        totals.setdefault(item, 0)
        totals[item] += prefs[other][item] * sim
        sim_sums.setdefault(item, 0)
        sim_sums[item] += sim



def getRecommandItems (prefs, item_match, user):
  """

  :param prefs:
  :param item_match:
  :param user:
  :return:
  """
  user_ratings = prefs[user]
  scores = {}
  total_sim = {}

  for (item, rating) in user_ratings.items():
    for (similartity, item2) in item_match[item]:
      if item2 in user_ratings: continue

      scores.setdefault(item2, 0)
      scores[item2] += similartity * rating

      total_sim.setdefault(item2, 0)
      scores[item2] += similartity

    rankings = [ (scores / total_sim[item], item) for item, scores in scores.items() ]

    rankings.sort()
    rankings.reverse()
    return rankings


if __name__ == "__main__":
  print getRecommandItems.__doc__
























