extern "C" {
    #include "./person_capi.h"
}

struct Person {
    static Person* New(const char* name, int age) {
        return (Person*)person_new((char*)name, age);
    }

    void Delete() {
        person_delete(person_handle_t(this));
    }

    void Set(char* name, int age) {
        person_set(person_handle_t(this), name, age);
    }
    char* GetName(char* buf, int size) {
        return person_get_name(person_handle_t(this), buf, size);
    }
    int GetAge() {
        return person_get_age(person_handle_t(this));
    }
};

//struct Person {
//    person_handle_t goobj_;
//
//    Person(const char* name, int age) {
//        this->goobj_ = person_new((char*)name, age);
//    }
//    ~Person() {
//        person_delete(this->goobj_);
//    }
//
//    void Set(char* name, int age) {
//        person_set(this->goobj_, name, age);
//    }
//    char* GetName(char* buf, int size) {
//        return person_get_name(this->goobj_ buf, size);
//    }
//    int GetAge() {
//        return person_get_age(this->goobj_);
//    }
//}