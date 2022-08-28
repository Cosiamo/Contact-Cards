import type { NextPage } from "next";
import Head from "next/head";
import { Tabs, TabList, Tab, Checkbox, Button, TabPanel, TabPanels, TextInput, Accordion, AccordionItem } from '@carbon/react';
import { string } from "zod";

export interface SectionProps {
  title: string;
  image: string;
  info: string;
}

const Home: NextPage = () => {
  return (
    <>
      <Head>
        <title>Create T3 App</title>
        <meta name="description" content="Generated by create-t3-app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
      <Tabs>
  <TabList
    aria-label="List of tabs"
    contained
  >
    <Tab>
      Tesla
    </Tab>
    <Tab>
      Lucid
    </Tab>
    <Tab>
      Rivian
    </Tab>
    <Tab>
      Fisker
    </Tab>
    <Tab>
      NIO
    </Tab>
  </TabList>
  <TabPanels>
    <TabPanel>
      <CarSection 
        title="One"
        image="/modelS2014.jpg"
        info="Testing the props"
      />
    </TabPanel>
    <TabPanel>
      <form
        style={{
          margin: '2em'
        }}
      >
        <legend className="cds--label">
          Validation example
        </legend>
        <Checkbox
          id="cb"
          labelText="Accept privacy policy"
        />
        <Button
          style={{
            marginBottom: '1rem',
            marginTop: '1rem'
          }}
          type="submit"
        >
          Submit
        </Button>
        <TextInput
          helperText="Optional help text"
          labelText="Text input label"
          type="text"
        />
      </form>
    </TabPanel>
    <TabPanel>
      Tab Panel 3
    </TabPanel>
    <TabPanel>
      Tab Panel 4
    </TabPanel>
    <TabPanel>
      Tab Panel 5
    </TabPanel>
  </TabPanels>
</Tabs>
      </main>
    </>
  );
};

const CarSection = ({title, image, info}: SectionProps) => {
  return (
    <Accordion align="start">
      <AccordionItem title={title}>
        <img src={image}/>
        {info}
      </AccordionItem>
    </Accordion>
  )
}

export default Home;