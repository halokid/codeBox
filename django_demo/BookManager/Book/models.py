from django.db import models

class BookInfo(models.Model):
  """图书模型类"""

  # 图书名称
  book_title = models.CharField(max_length = 20)
  # 图书日期
  book_date = models.DateField()

  def __str__(self):
    return self.book_title


class People(models.Model):
  """任务模型类"""
  name = models.CharField(max_length = 20)

  age = models.IntegerField()
    