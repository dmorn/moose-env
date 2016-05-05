
char a[250];
snprintf(a,250,"SELECT * FROM `group_items` WHERE group_id='%s' AND stock_id =%d AND category_id IN (SELECT category_id FROM category where isSubcategoryOf(category_id,%d)=1)",group_id,item_id,category_id);