For this Project, we are following Danial Burke youtube channel video where he demonstrates how to implement local RAG pipeline. He also publih a git repo for the same topic.

However, we are going to twist the project a little bit. We will crawl htmx-book and create a chatbot that implements RAG pipeline with the htmx-book context.

This project aims to be a soft introduction for me towards Local LLM and RAG implementation.

1. crawl htmx url to ge htmx resources (other topic is also fine)
2. process resources into embeddings and store somewhere (torch.tensor is fine | vector db)
3. Rank resource by query simmilarity (dot prodcut or cosine simmilarity)
4. retrieve relevant resources and augment with query
5. query

## Crawling htmx-book

- we will use colly library in golang to retrieve our data.
- I just felt like using golang.

first-version json data
c.OnHTML("body", func(h \*colly.HTMLElement) {

    i := Item{
    Chapter: h.ChildText("header h1"),
    Content: h.ChildText("main div.division-content p "),
    // Name: h.ChildAttr("h3 a", "title"),
    // Price: h.ChildText("p.price_color"),
    }

            items = append(items, i)

        })

second-version json data
c.OnHTML("body", func(h \*colly.HTMLElement) {

    	contents := []Paragraph{}

    	h.ForEach("main p", func(i int, h *colly.HTMLElement) {
    		p := Paragraph{
    			Content: h.Text,
    		}
    		contents = append(contents, p)
    	})

    	i := Item{
    		Chapter:    h.ChildText("header h1"),
    		Paragraphs: contents,
    		// Name:    h.ChildAttr("h3 a", "title"),
    		// Price:   h.ChildText("p.price_color"),
    	}

    	items = append(items, i)
    })

we will proceed the project with 2nd version data for now.

> NOTE: that there is a lot more improvement could be made.

## Retrieval

### Preprocessing the htmx-book

1. We process the one super long string of text into multiple small chunk of text per chapter.
2. We need to adjust the size of chunk relative to our embedding models. Read the `main.ipynb` to see details on token size for our embedding model.
3. The embedding model we use is `all-mpnet-base-v2`.
   > NOTE: There is a lot more thinks happens in this steps but we will skip the explanation.

### Store the embedding vector

- We decide to save it only in pandas DataFrame as the size of the context vector is still relatively small.
  > NOTE: We save the data to json file beforehand as we could not continue in our local environment due to resource issues
  > NOTE: We will continue in google collab with free tier t4 runtime.
  > NOTE: I know we can use the same file but to make me see the flow of my project clearly we move to another file which is `htmx-helper.ipynb`

1. We store the processed data into json file `second-version.json`
2. We load the file an apply our embedding model.
3. We saved it into a pandas DF

### Retrieve relevant context for query

1. For this part we use cosine simmilarity function to compare the simmilarity of vector between our query and the htmx-book embedded data
2. We choose top 5 most simmilar data as our context with topk function.

## Augment

> NOTE: This part actually finished last.

1. We make a prompting template to simplify our flow.
2. As the model we use came with a useful chat template, we also implemented it
3. Here there is a lot of improvement could be made based on the prompting styles but we choose to settle down on a more simpler prompt.

## Generation

- Check local resource capability on which model to implement.
- VRAM plays a crucial role to load The model fully into the GPU.

1. We choose to use `Gemma-2-2b-it` \* It is a Gemma 2 model with 2billion parameters and are tuned to followed instructions.
   > NOTE: We have a hard time in loading the model on our first try. Remember to check bitsandbytes, accelerates library and always restart the session after a new pip download
