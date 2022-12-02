local k8s = import 'k8s.jsonnet';

local ecr_prefix = '562055475000.dkr.ecr.ap-northeast-1.amazonaws.com/qor5/';

local images = [
  { type: 'deployment', name: 'docs', image: ecr_prefix + 'docs' },
];

k8s.set_images(
  namespace='qor5-test',
  images=images
)
