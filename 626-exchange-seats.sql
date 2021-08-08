-- Link: https://leetcode.com/problems/exchange-seats/

select
case
    when s.id % 2 <> 0 and s.id = (select max(id) from seat) then s.id
    when s.id % 2 = 0 then s.id - 1
    else s.id + 1
end as id, s.student
from seat s
order by id ASC