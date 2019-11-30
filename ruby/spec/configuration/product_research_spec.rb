describe Configuration::ProductResearch do
  let(:marketplace_codes) { ENV['MARKETPLACES'].to_s.split(',') }

  context 'Valid marketplaces' do
    it 'has all marketplaces in MARKETPLACES constant' do
      expect(Configuration::ProductResearch::MARKETPLACES.count).to eq(marketplace_codes.size)
    end

    it 'has all' do
      expect(marketplace_codes).to_not be_empty
      marketplace_codes.each do |code|
        expect(Configuration::ProductResearch.find_marketplace_by_code(code)).to_not be_nil
      end
    end

    Configuration::ProductResearch::MARKETPLACES.each do |m|
      it "has #{m.name}" do
        expect(m.name).to_not be_empty
        expect(m.code).to_not be_empty
        expect(m.url).to_not be_empty
        expect(m.categories).to_not be_empty
      end
    end
  end

  context 'Valid categories' do
    let(:marketplace_category) { ENV['FIND_CATEGORY'].to_s.split('|') }
    let(:marketplace) { marketplace_category.first }
    let(:name) { marketplace_category.last }

    Configuration::ProductResearch::MARKETPLACES.each do |m|
      m.categories.each do |c|
        it "has #{c.name}" do
          expect(c.name).to_not be_empty
          expect(c.type).to_not be_empty
          expect(c.urls).to_not be_empty
        end
      end
    end

    it 'finds specific category' do
      expect(marketplace).to_not be_nil
      expect(name).to_not be_nil

      c = Configuration::ProductResearch.find_marketplace_by_code(marketplace).find_catergory_by_name(name)
      expect(c).to_not be_nil
    end
  end
end
