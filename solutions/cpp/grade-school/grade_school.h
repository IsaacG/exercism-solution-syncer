#pragma once
#include <map>
#include <string>
#include <vector>

namespace grade_school {

class school {
    std::map<int, std::vector<std::string>> m_roster;

    public:
        const std::map<int, std::vector<std::string>> roster() const;
        void add(std::string name, int grade);
        const std::vector<std::string> grade(int requested_grade) const;
};

}  // namespace grade_school
