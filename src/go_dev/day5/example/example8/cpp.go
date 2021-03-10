class Student {
	private:
	   int score;
	   string name;
 
	public:
	   void init(string name, int score);   

};

void Student::init(Student *this, string name, int score) {
	 this->name = name
	 this->score = score 
}

func (this *Student) init() 