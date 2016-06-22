
// PHP5
struct _zval_struct {
  union {
    long      lval;
    double    dval;
    
    struct {        //字符串类型结构体
      char *val;
      int len;
    } str;
    
    HashTable *ht;        //hash
    
    zend_object_value   obj;
    zend_ast            *ast;
  } value;
  
  zend_uint     refcount__gc;
  zend_uchar    type;
  zend_uchar    is_ref__gc;
}



//PHP7
struct _zval_struct {
  union {
    long                  lval;
    double                dval;
    zend_refcounted       *counted;
    zend_string           *str;
    zend_array            *arr;
    zend_object           *obj;
    zend_resource         *res;
    zend_ast_ref          *ast;
    aval                  *zv;
    void                  *ptr;
    zend_class_entry      *ce;
    zend_function         *func; 
  } value;
  
  union {
    struct {
      ZEND_ENDIAN_LOHI_4 ( 
        zend_uchar        type;
        zend_uchar        type_flags;
        zend_uchar        const_flags;
        zend_uchar        reserved;
        )
      } v;
      zend_unit   type_info;
  } u1;
  
  
  union {
    zend_uint         var_flags;
    zend_uint         next;
    zend_uint         str_offset;
    zend_uint         cache_slot;
  } u2;
  
}


// PHP7
struct _zend_string {
  zend_refcounted         gc;
  zed_ulong               h;
  size_t                  len;
  char                    val[1]
};










