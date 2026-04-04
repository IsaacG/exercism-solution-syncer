#include <algorithm>
#include "grade_school.h"

using namespace std;

namespace grade_school {

const map<int, vector<string>> school::roster() const {
    return m_roster;
}

void school::add(string name, int grade) {
    m_roster[grade].push_back(name);
    sort(m_roster[grade].begin(), m_roster[grade].end());
}

const vector<string> school::grade(int requested_grade) const {
    return m_roster.find(requested_grade)->second;
}

}  // namespace grade_school

