  
public function getUser($uid)
{
"SELECT user_id, username, email, name, surname, balance, type, group_id FROM user WHERE user_id=".$uid;
}

public function getCategoryObjects($cid){ //get all objects from this category and its subcategories
"SELECT * FROM object WHERE category_id IN (SELECT category_id FROM category where isSubcategoryOf(category_id,".$cid.")=1)";
}
  
public function getMainCategories(){	//get main categories
"SELECT * FROM category where parent_id IS NULL";
}