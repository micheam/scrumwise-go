(ns user
  (:require [clojure.repl :as repl]
            [clojure.pprint :refer [pprint]]
            [net.cgrand.enlive-html :as html]
            [clojure.string :refer [starts-with? ends-with? trim ] :as string]
            [camel-snake-kebab.core :refer [->kebab-case]]))

(def url "https://www.scrumwise.com/api.html")
(def content (html/html-resource (java.net.URL. url)))
(def method-list
  (html/select content [#{[:a (html/attr-ends :id "-methods")]
                          [:a (html/attr-ends :id "-method")]
                          [:h3 (html/but (html/has [:a]))]}]))

(defn method-group?  "element が method group 要素か否か" [elm]
  (if-let [id (-> elm (get-in [:attrs :id]))]
    (ends-with? id "-methods")
    false))

(defn transform

  ([col]
   (transform {} nil col))

  ([m k col]
   (let [head (first col)]
     (cond
       (empty? col) m

       (method-group? head)
       (let [id (-> head :attrs :id keyword)]
         (transform (assoc m id []) id (rest col)))

       :else
       (let [content (-> head :content first trim (string/replace "\n" ""))
             m' (update m k conj content)]
         (transform m' k (rest col)))))))

(defn markdown-list
  [k vs]
  (do
    (println "")
    (println (format "- [ ] [%s](%s#%s)"
                     (string/replace (->kebab-case (name k)) "-" " ") url (name k)))
    (doseq [x vs] 
      (println (format "  - [ ] [%s](%s#%s-method)" (->kebab-case x) url x)))))


(doseq [[k v] (-> method-list transform)]
  (markdown-list k v))
