#### Custom configuration

mkdir data
cp cayley_example.yml data/cayley.yml
cp data/testdata.nq data/my_data.nq
# initialize and serve database
docker run -v $PWD/data:/data -p 64210:64210 -d cayleygraph/cayley -c /data/cayley.yml --init -i /data/my_data.nq
# serve existing database
docker run -v $PWD/data:/data -p 64210:64210 -d cayleygraph/cayley -c /data/cayley.yml


docker run  -v $PWD/cfg:/cfg -v $PWD/data:/data -p 64210:64210 -d quay.io/cayleygraph/cayley -c /data/cayley.yml --init -i /data/my_data.nq