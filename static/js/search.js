var index = elasticlunr(function () {
  this.addField('name');
  this.addField('fees');
  this.addField('address');
  this.addField('categories');
  this.addField('description');
  this.addField('services');
  this.addField('languages');
  this.addField('categories');
  this.addField('pops_served');
});

for (var i = 0; i < resources.length; i++) {
    resources[i].id = i;
    index.addDoc(resources[i]);
}

console.log(index.search("Berkeley"));