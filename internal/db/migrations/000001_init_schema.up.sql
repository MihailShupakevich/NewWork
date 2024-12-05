CREATE TABLE users (
id SERIAL PRIMARY KEY,
email VARCHAR(100) UNIQUE NOT NULL,
password VARCHAR(100) NOT NULL
);

CREATE TABLE courses (
 course_id SERIAL PRIMARY KEY,
 courses_execution_time TIMESTAMP,
 courses_name VARCHAR(100) UNIQUE NOT NULL,
 course_specification_id INT,
 FOREIGN KEY (course_specification_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE TABLE sections(
 section_id SERIAL PRIMARY KEY,
 section_execution_time TIMESTAMP,
 section_name VARCHAR(100) UNIQUE NOT NULL,
 section_specification_id INT,
 FOREIGN KEY (section_specification_id) REFERENCES users(id) ON DELETE CASCADE

);
CREATE TABLE list_of_selected_sections (
 list_of_selected_sections_id SERIAL PRIMARY KEY,
 user_id INTEGER,
 foreign key (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE TABLE list_of_selected_courses(
    list_of_selected_courses_id serial primary key ,
    user_id INTEGER,
    foreign key (user_id) REFERENCES users(id) ON DELETE CASCADE

);

CREATE TABLE section_specification(
        section_specification_id serial primary key ,
        section_specification_name varchar(256)

);
CREATE TABLE course_specification(
                                      course_specification_id serial primary key ,
                                      course_specification_name varchar(256)
);

CREATE TABLE selected_courses (
                                   list_of_selected_courses_id INTEGER,
                                   foreign key  (list_of_selected_courses_id)  REFERENCES list_of_selected_courses(list_of_selected_courses_id) ON DELETE CASCADE,
                                   course_id INTEGER, foreign key (course_id) REFERENCES courses(course_id) ON DELETE CASCADE
);
CREATE TABLE selected_sections (
                                   list_of_selected_sections_id INTEGER, foreign key (list_of_selected_sections_id) REFERENCES list_of_selected_sections(list_of_selected_sections_id) ON DELETE CASCADE,
                                   section_id INTEGER, foreign key (section_id) REFERENCES sections(section_id) ON DELETE CASCADE
);
CREATE TABLE topics_for_courses(
    topics_for_courses_id serial primary key ,
    user_id INTEGER, foreign key (user_id) REFERENCES users(id) ON DELETE CASCADE,
    topics_for_courses_name varchar(256) NOT NULL
);
CREATE TABLE additional_material(
    additional_material_title varchar(256) not null ,
    additional_material text not null ,
    topics_for_courses_id INTEGER, foreign key (topics_for_courses_id) references topics_for_courses(topics_for_courses_id) ON DELETE CASCADE
);
CREATE TABLE  individual_works(
    work_title varchar(100) not null ,
    work text not null ,
    section_id integer,
    user_id integer,
    foreign key (user_id) references users(id) on delete cascade ,
    foreign key (section_id) references sections(section_id) on delete cascade
);
CREATE TABLE competitions(
    competitions_id serial primary key ,
    section_id int,
    date timestamp,
    foreign key (section_id) references sections(section_id) on delete cascade
);
CREATE TABLE competitions_users(
                             competitions_id serial primary key ,
                             section_id int,
                             foreign key (section_id) references sections(section_id) on delete cascade ,
                             foreign key (competitions_id) references competitions(competitions_id) on delete cascade
);
CREATE TABLE teachers(
                             teacher_id serial primary key ,
                             section_id int,
                             teacher_fio varchar not null unique ,
                             foreign key (section_id) references sections(section_id) on delete cascade
);
create table timetable_of_classes(
                                     timetable_of_classes_id serial primary key ,
                                     section_id int,
                                        date timestamp,
    teacher_id int,
                                     foreign key (teacher_id) references teachers(teacher_id) on delete cascade ,
                                     foreign key (section_id) references sections(section_id) on delete cascade
);