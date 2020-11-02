(ns user
  (:require [clojure.repl :as repl]
            [clojure.pprint :refer [pprint]]
            [net.cgrand.enlive-html :as html]))

(def url
  "https://www.scrumwise.com/api.html")

(def content
  (html/html-resource (java.net.URL. url)))

(for
  [x (-> content (html/select [:a.title-link]))]
  (pprint (keys x)))

(def greet "hello")
