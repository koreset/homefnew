package main

import (
	"fmt"
	"github.com/koreset/homefnew/app/utils"
)

var message = ` "<p class="MsoNormal" style="text-align: justify;"><span style="mso-bidi-font-family: Calibri;" lang="EN-GB">While Nigerians face uncertainties over food security due to incessant herders-farmers clashes, another threat is dawning on the nation without much notice. The fact that President Muhammadu Buhari just inaugurated a Food Security Council underscores the centrality of food security to the country. However, without food safety there cannot be food security.</span></p><p class="MsoNormal" style="text-align: justify;"><span style="mso-bidi-font-family: Calibri;" lang="EN-GB">On the 22<sup>nd</sup> of March, 2018, the country coordinator for Open Forum on Agricultural Biotechnology (OFAB) <span style="mso-bidi-font-style: italic;">and National Biotechnology Development Agency (NABDA) chieftain, </span>Dr Rose Gidado, stated in an interview</span><a name="_ednref1" href="#_edn1" title="" style="mso-endnote-id: edn1;"><span class="MsoEndnoteReference"><span lang="EN-GB"><span style="mso-special-character: footnote;"><span class="MsoEndnoteReference"><span style="font-size: 11.0pt; line-height: 115%; font-family: 'Calibri',sans-serif; mso-fareast-font-family: Calibri; mso-bidi-font-family: 'Times New Roman'; mso-ansi-language: EN-GB; mso-fareast-language: EN-US; mso-bidi-language: AR-SA;" lang="EN-GB">[1]</span></span></span></span></span></a><span style="mso-bidi-font-family: Calibri;" lang="EN-GB"> with newsmen in Abuja that “<em style="mso-bidi-font-style: normal;">before the end of 2018, Bt Cowpea and Bt Cotton as biotechnology products in Nigeria would be in market</em>.” </span></p><p class="MsoNormal" style="text-align: justify;"><span style="mso-bidi-font-family: Calibri;" lang="EN-GB">Health of Mother Earth Foundation (HOMEF) raised concerns with regard to this announcement and stated that these products are soon to be unleashed on an unsuspecting Nigerian public. According to the organisation, there is no guarantee about the environmental and health safety of the beans and cotton to be released by the end of 2018. </span></p><p class="MsoNormal" style="text-align: justify;"><span style="mso-bidi-font-family: Calibri;" lang="EN-GB">Nnimmo Bassey, the Director of HOMEF stressed that there are <span style="mso-bidi-font-style: italic;">serious challenges</span></span><a name="_ednref2" href="#_edn2" title="" style="mso-endnote-id: edn2;"><span class="MsoEndnoteReference"><span lang="EN-GB"><span style="mso-special-character: footnote;"><span class="MsoEndnoteReference"><span style="font-size: 11.0pt; line-height: 115%; font-family: 'Calibri',sans-serif; mso-fareast-font-family: Calibri; mso-bidi-font-family: 'Times New Roman'; mso-ansi-language: EN-GB; mso-fareast-language: EN-US; mso-bidi-language: AR-SA;" lang="EN-GB">[2]</span></span></span></span></span></a><span style="mso-bidi-font-family: Calibri; mso-bidi-font-style: italic;" lang="EN-GB"> GMOs pose in the areas of toxicology, allergy and immune dysfunction and genetic disorders which make it very important that Nigeria adopts the precautionary principle-which warns that </span><span style="mso-bidi-font-family: Calibri;" lang="EN-GB">strict measures should be applied where there are threats of serious or irreversible damage and lack of full scientific certainty </span></p><p class="MsoNormal" style="text-align: justify;"><span style="mso-bidi-font-family: Calibri; mso-bidi-font-style: italic;" lang="EN-GB">The OFAB/NABDA officer stated that “<em>the pesticide tolerant cowpea will allow for a reduced use of pesticide</em>”. To this, Nnimmo Bassey, responded and stated it is well known that pests have developed resistance</span><a name="_ednref3" href="#_edn3" title="" style="mso-endnote-id: edn3;"><span class="MsoEndnoteReference"><span lang="EN-GB"><span style="mso-special-character: footnote;"><span class="MsoEndnoteReference"><span style="font-size: 11.0pt; line-height: 115%; font-family: 'Calibri',sans-serif; mso-fareast-font-family: Calibri; mso-bidi-font-family: 'Times New Roman'; mso-ansi-language: EN-GB; mso-fareast-language: EN-US; mso-bidi-language: AR-SA;" lang="EN-GB">[3]</span></span></span></span></span></a><span style="mso-bidi-font-family: Calibri; mso-bidi-font-style: italic;" lang="EN-GB"> to the Bt toxins and this resistance leads to an increased use of toxic chemicals which increases the damage to the environment. </span></p><p class="MsoNormal" style="text-align: justify;"><span style="mso-bidi-font-family: Calibri; mso-bidi-font-style: italic;" lang="EN-GB">According to Mariann Orovwuje, the African Food Sovereignty Alliance (AFSA) chairperson, “Bt crops and other GM foods will not help Nigeria’s economy as the supporters of the technology allege. Rather, there will be forced dependence on corporate bodies for seeds. Farmers will have no right to reuse their seeds and agricultural production will be left in the hands of large scale industrial investors.” </span></p><p class="MsoNormal" style="text-align: justify;"><span style="mso-bidi-font-family: Calibri; mso-bidi-font-style: italic;" lang="EN-GB">The National Food, Drug Administration and Control (NAFDAC) in June 2017 revealed that 24 export products from Nigeria were rejected by the European Union</span><a name="_ednref4" href="#_edn4" title="" style="mso-endnote-id: edn4;"><span class="MsoEndnoteReference"><span lang="EN-GB"><span style="mso-special-character: footnote;"><span class="MsoEndnoteReference"><span style="font-size: 11.0pt; line-height: 115%; font-family: 'Calibri',sans-serif; mso-fareast-font-family: Calibri; mso-bidi-font-family: 'Times New Roman'; mso-ansi-language: EN-GB; mso-fareast-language: EN-US; mso-bidi-language: AR-SA;" lang="EN-GB">[4]</span></span></span></span></span></a><span style="mso-bidi-font-family: Calibri; mso-bidi-font-style: italic;" lang="EN-GB"> (EU) in the year 2016 for failing to meet standards. HOMEF warns that the market for genetically modified products from Nigeria is narrowing instead of expanding, with the strict requirement of the EU. </span></p><p class="MsoNormal" style="text-align: justify;"><span style="mso-bidi-font-family: Calibri; mso-bidi-font-style: italic;" lang="EN-GB">HOMEF believes that the way to improve economic situation for farmers is to invest in organic agriculture, provide farmers with extension services, needed infrastructure, good roads and access to land and loans. Support for farmers should include investment in research and exploration of agroecology approach to the problems of pests and diseases.</span></p><p class="MsoNormal" style="text-align: justify;"><span style="mso-bidi-font-family: Calibri; mso-bidi-font-style: italic;" lang="EN-GB">The organisation stated that it is good to learn from others who have taken caution against GMOs. “The Ugandan President, Yoweri Museveni had declined to sign into law a Biosafety Bill passed by the Ugandan Parliament in October 2017 because of issues that included liability and redress and concerns on conservation of indigenous crops and agricultural biodiversity.”</span></p><p class="MsoNormal" style="text-align: justify;"><span style="mso-bidi-font-family: Calibri;" lang="EN-GB">According to Joyce Ebebeinwe, Biosafety Officer at HOMEF, the Director of the National Biosafety Management Agency (NBMA) during a recent twitter chat was asked how the <em style="mso-bidi-font-style: normal;">akara</em> or <em style="mso-bidi-font-style: normal;">moi-moi</em> made from GM beans and other foods sold by the road sides will be labelled and his response was that the sellers should erect signposts announcing that they are selling GM&nbsp;<span style="mso-bidi-font-style: italic;">products</span>. That would be laughable if not for the fact that it is an extremely tragic notion. </span></p><p class="MsoNormal" style="text-align: justify;"><span style="mso-bidi-font-family: Calibri;" lang="EN-GB">HOMEF insisted that the promise to have GMOs labelled in Nigeria to ensure that the public have a choice on whether or not to eat such crops will not work mainly due to our socio-cultural and economic realities. </span></p><p class="MsoNormal" style="text-align: justify;"><span style="mso-bidi-font-family: Calibri;" lang="EN-GB">Furthermore, the recent statement by the NBMA DG, Dr Rufus Ebegba, at a media conference in Abuja on the 5<sup>th</sup> of April shows clearly how flawed our biosafety regulatory system is. According to reports</span><a name="_ednref5" href="#_edn5" title="" style="mso-endnote-id: edn5;"><span class="MsoEndnoteReference"><span lang="EN-GB"><span style="mso-special-character: footnote;"><span class="MsoEndnoteReference"><span style="font-size: 11.0pt; line-height: 115%; font-family: 'Calibri',sans-serif; mso-fareast-font-family: Calibri; mso-bidi-font-family: 'Times New Roman'; mso-ansi-language: EN-GB; mso-fareast-language: EN-US; mso-bidi-language: AR-SA;" lang="EN-GB">[5]</span></span></span></span></span></a><span style="mso-bidi-font-family: Calibri;" lang="EN-GB">, “<em style="mso-bidi-font-style: normal;">Dr. Rufus Ebegba, has given importers of genetically modified (GM) seeds a seven-day ultimatum to formalise their dealings or risk being shut down</em>”. </span></p><p class="MsoNormal" style="text-align: justify; mso-pagination: none; mso-layout-grid-align: none; text-autospace: none;"><span style="mso-bidi-font-family: Calibri;" lang="EN-GB">HOMEF expressed that this statement is in direct contrast to the provisions of the law. <strong style="mso-bidi-font-weight: normal;"><span style="color: black;">Section 23 (1) of the NBMA Act 2015 </span></strong><span style="color: black;">states: “<em style="mso-bidi-font-style: normal;">Any person, institution or body who wishes to import, export, transit or otherwise carry out a contained field trial, multi-locational trial or commercial release of a genetically modified organism shall apply to the Director General of the Agency <strong style="mso-bidi-font-weight: normal;"><span style="text-decoration: underline;">not less than 270 days to the date of import, export, transit or the commencement of such activity”.</span></strong></em></span></span></p><p class="MsoNormal" style="text-align: justify; mso-pagination: none; mso-layout-grid-align: none; text-autospace: none;"><span style="mso-bidi-font-family: Calibri; color: black;" lang="EN-GB">According to HOMEF, the statement by the NBMA DG may be construed to mean that dealers on GM products in Nigeria will be given permits after they had imported GMO seeds without passing through due approval processes. HOMEF totally objects to any sort of formalisation of illegal importation of GMOs into the country. A short notice to import these crops does not allow for risk assessments or safety assurances and regrettably this may have been the basis on which the agency granted WACOT Ltd permits to import several varieties of GM maize over a period of three years – after the company was alleged to have illegally brought in the crop in September 2017 and had been ordered to repatriate the grains.</span></p><p class="MsoNormal" style="text-align: justify;"><span style="mso-bidi-font-family: Calibri;" lang="EN-GB">The Ecological Think Tank calls on the Nigerian government to look critically at the activities of the Nigerian Biosafety Regulatory Agency and the subject of genetically modified foods in the country. While HOMEF commends the action of the federal government in establishing the food security council, the group insists that it is pertinent that action is taken to ensure that our foods are indeed safe.<span style="mso-spacerun: yes;">&nbsp; </span></span></p><p>&nbsp;</p><p>Notes</p><div style="mso-element: endnote-list;"><hr width="33%" size="1" align="left"><div id="edn1" style="mso-element: endnote;"><p class="MsoEndnoteText"><a name="_edn1" href="#_ednref1" title="" style="mso-endnote-id: edn1;"><span class="MsoEndnoteReference"><span style="mso-special-character: footnote;"><span class="MsoEndnoteReference"><span style="font-size: 10.0pt; line-height: 115%; font-family: 'Calibri',sans-serif; mso-fareast-font-family: Calibri; mso-bidi-font-family: 'Times New Roman'; mso-ansi-language: EN-US; mso-fareast-language: EN-US; mso-bidi-language: AR-SA;">[1]</span></span></span></span></a> Vanguard News: <a href="https://www.vanguardngr.com/2018/03/nigeria-use-biotechnology-improve-crops/">https://www.vanguardngr.com/2018/03/nigeria-use-biotechnology-improve-crops/</a></p></div><div id="edn2" style="mso-element: endnote;"><p class="MsoEndnoteText"><a name="_edn2" href="#_ednref2" title="" style="mso-endnote-id: edn2;"><span class="MsoEndnoteReference"><span style="mso-special-character: footnote;"><span class="MsoEndnoteReference"><span style="font-size: 10.0pt; line-height: 115%; font-family: 'Calibri',sans-serif; mso-fareast-font-family: Calibri; mso-bidi-font-family: 'Times New Roman'; mso-ansi-language: EN-US; mso-fareast-language: EN-US; mso-bidi-language: AR-SA;">[2]</span></span></span></span></a> <span style="mso-ansi-language: EN-GB;" lang="EN-GB">Science in Society Archive: </span><a href="http://www.i-sis.org.uk/TheCaseforAGM-FreeSustainableWorld.php"><span style="mso-ansi-language: EN-GB;" lang="EN-GB">www.i-sis.org.uk/TheCaseforAGM-FreeSustainableWorld.php</span></a></p></div><div id="edn3" style="mso-element: endnote;"><p class="MsoEndnoteText"><a name="_edn3" href="#_ednref3" title="" style="mso-endnote-id: edn3;"><span class="MsoEndnoteReference"><span style="mso-special-character: footnote;"><span class="MsoEndnoteReference"><span style="font-size: 10.0pt; line-height: 115%; font-family: 'Calibri',sans-serif; mso-fareast-font-family: Calibri; mso-bidi-font-family: 'Times New Roman'; mso-ansi-language: EN-US; mso-fareast-language: EN-US; mso-bidi-language: AR-SA;">[3]</span></span></span></span></a> PLOS Research Journal: <a href="http://journals.plos.org/plosone/article?id=10.1371/journal.pone.0022629">http://journals.plos.org/plosone/article?id=10.1371/journal.pone.0022629</a></p></div><div id="edn4" style="mso-element: endnote;"><p class="MsoEndnoteText"><a name="_edn4" href="#_ednref4" title="" style="mso-endnote-id: edn4;"><span class="MsoEndnoteReference"><span style="mso-special-character: footnote;"><span class="MsoEndnoteReference"><span style="font-size: 10.0pt; line-height: 115%; font-family: 'Calibri',sans-serif; mso-fareast-font-family: Calibri; mso-bidi-font-family: 'Times New Roman'; mso-ansi-language: EN-US; mso-fareast-language: EN-US; mso-bidi-language: AR-SA;">[4]</span></span></span></span></a>Vanguard News: <a href="https://www.vanguardngr.com/2017/06/nafdac-says-eu-rejected-24-nigeria-exported-products-2016/">https://www.vanguardngr.com/2017/06/nafdac-says-eu-rejected-24-nigeria-exported-products-2016/</a><span style="mso-spacerun: yes;">&nbsp; </span></p></div><div id="edn5" style="mso-element: endnote;"><p class="MsoEndnoteText"><a name="_edn5" href="#_ednref5" title="" style="mso-endnote-id: edn5;"><span class="MsoEndnoteReference"><span style="mso-special-character: footnote;"><span class="MsoEndnoteReference"><span style="font-size: 10.0pt; line-height: 115%; font-family: 'Calibri',sans-serif; mso-fareast-font-family: Calibri; mso-bidi-font-family: 'Times New Roman'; mso-ansi-language: EN-US; mso-fareast-language: EN-US; mso-bidi-language: AR-SA;">[5]</span></span></span></span></a> EnviroNews Nigeria: <a href="http://www.environewsnigeria.com/gm-seeds-importers-given-seven-day-ultimatum-to-formalise-dealings/">www.environewsnigeria.com/gm-seeds-importers-given-seven-day-ultimatum-to-formalise-dealings/</a></p></div></div><p>&nbsp;</p><p>&nbsp;</p>"`
var (
	ignoreTags = []string{"title", "script", "style", "iframe", "frame", "frameset", "noframes", "noembed", "embed", "applet", "object", "base"}

	defaultTags = []string{"h1", "h2", "h3", "h4", "h5", "h6", "div", "hr", "span", "p", "br", "b", "i", "strong", "em", "ol", "ul", "li", "a", "img", "pre", "code", "blockquote", "article", "section"}

	defaultAttributes = []string{"id", "class", "src", "href", "title", "alt", "name", "rel"}
)


func main() {

	newmessage := utils.RemoveAllTags(message)

	fmt.Println(newmessage)
}
