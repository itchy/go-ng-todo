/* App Module */
(function(){
  var app = angular.module('todoApp', []);
  
  app.controller('ListController',['$http', function($http){
    var list = this;
    this.tasks = [];

    this.updateTask = function(task){
      console.log(task.done);
      task.done = !task.done;
      this.update(task);
    };

    this.update = function(task){
      var state = task.done ? 1 : 0
      $http.put('/tasks/' + task.id + '.json', {task: {id: task.id, state: state}}).success(function(data){
        console.log('updated task ' + task.id);
      });
    };

    this.getTotalTasks = function(){
      // var activeTasks = this.tasks.filter(function(x) { return !x.done ; });
      var activeTasks = _.select(this.tasks, function(x) { return !x.done ; });
      return activeTasks.length;
    };

    $http.get('/tasks.json').success(function(data){
      list.tasks = data;
    });
  }]);

  app.controller('TaskController',['$http', function($http){
    var task = this;
    this.item = {};

    this.addTask = function(list){
      $http.post('/tasks.json', {task: task.item}).success(function(data){
        list.tasks.push(task.item);
        task.item = {};
      });
    };

  }]);

})();