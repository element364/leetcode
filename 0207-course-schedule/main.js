function canFinish(numCourses, prerequisites) {
  const visited = [];
  const recStack = [];

  function dfs(current) {
    if (!visited[current]) {
      visited[current] = true;
      recStack[current] = true;

      for (const [from, to] of prerequisites) {
        if (from == current) {
          if (!visited[to] && dfs(to)) {
            return true;
          } else if (recStack[to]) {
            return true;
          }
        }
      }
    }

    recStack[current] = false;
    return false;
  }

  for (let i = 0; i < numCourses; i++) {
    if (!visited[i] && dfs(i)) {
      return false;
    }
  }

  return true;
}

let numCourses = 2;
let prerequisites = [[1, 0]];
console.log("Example 1", canFinish(numCourses, prerequisites));

numCourses = 2;
prerequisites = [
  [1, 0],
  [0, 1],
];
console.log("Example 2", canFinish(numCourses, prerequisites));
