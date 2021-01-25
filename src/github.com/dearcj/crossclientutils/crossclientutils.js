;(function() {
  var pkg = {};
  pkg["github.com\dearcj\crossclientutils"] = (function() {
    U.prototype.GetSkillProb = function(baseChance, will, maxwill) {
      var z = this;
      return baseChance * 1 - will / maxwill * 0.75;
    };
    function U (o) {
      o = o || {}
    };
    function main () {
      var y = new U();
      var z = y;
      y = z;
      z.GetSkillProb.bind(z)(1, 1, 1)
    };
    return {
      U: U,
      main: main
    };
  })();
  return pkg["github.com\dearcj\crossclientutils"].main();
})()
