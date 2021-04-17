from django.contrib import admin
from Book.models import BookInfo

# Register your models here.

# 自定义模型管理类
class BookInfoAdmin(admin.ModelAdmin):
  list_display = ['id', 'book_title', 'book_date']

# 注册模型类
admin.site.register(BookInfo, BookInfoAdmin)