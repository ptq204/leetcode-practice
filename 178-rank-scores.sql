-- Link: https://leetcode.com/problems/rank-scores/

select s.Score, (select count(*)
     from (select distinct `Score` from `Scores`) s2
     where s2.Score >= s.Score
    ) as `Rank`
from `Scores` s
order by s.Score desc