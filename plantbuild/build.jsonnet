local c = import 'dc.jsonnet';
local dc = c {
  dockerRegistry: 'gcr.io',
};

dc.build_apps_image('sunfmin/sunfmin', [
  {name: 'qor5-docs', dockerfile: './docs/Dockerfile', context: '../'},
])
