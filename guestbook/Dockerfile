# Copyright 2016 The Kubernetes Authors All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.7

RUN  rm -rf /go/src/guestbook
ADD . /go/src/guestbook
WORKDIR /go/src

RUN cd guestbook/ && go get -v && go build -o ../bin/guestbook
RUN cp guestbook/guestbook/Dockerfile .

CMD tar czf - .
