module Configuration
  module ProductResearch
    class Marketplace
      attr_reader :name, :code, :url, :locale, :region, :categories, :mail_packages

      def initialize(options)
        @name = options.fetch(:name)
        @code = options.fetch(:code)
        @url = options.fetch(:url)
        @locale = options.fetch(:locale)
        @region = options.fetch(:region)
        @categories = options.fetch(:categories).map { |c| Category.new(c) }
        @mail_packages = options.fetch(:mail_packages).map { |m| MailPackage.new(m) }
      end

      def find_mail_package_by_name(name)
        mail_packages.find { |m| m.name == name }
      end

      def find_catergory_by_name(name)
        categories.find { |c| c.name == name }
      end
    end
  end
end
