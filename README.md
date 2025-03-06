1. crawl htmx url to ge htmx resources (other topic is also fine)
2. process resources into embeddings and store somewhere (torch.tensor is fine | vector db)
3. Rank resource by query simmilarity (dot prodcut or cosine simmilarity)
4. retrieve relevant resources and augment with query
5. query

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

we will proceed the project with 2nd version data for now
