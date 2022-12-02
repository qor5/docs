local c = import 'dc.jsonnet';
local dc = c {
  dockerRegistry: 'public.ecr.aws/a6b5c3b2',
};

dc.build_apps_image('qor5/qor5', [
  { name: 'docs', dockerfile: './Dockerfile', context: '.' },
])
